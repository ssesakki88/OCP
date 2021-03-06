package networking

import (
	"context"
	"net"
	"time"

	"k8s.io/client-go/rest"

	"github.com/apparentlymart/go-cidr/cidr"
	g "github.com/onsi/ginkgo"
	o "github.com/onsi/gomega"
	exutil "github.com/openshift/origin/test/extended/util"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/fields"
	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/apiserver/pkg/authentication/serviceaccount"
	"k8s.io/client-go/kubernetes"
)

var _ = g.Describe("[sig-network][endpoints] admission", func() {
	defer g.GinkgoRecover()
	oc := exutil.NewCLI("endpoint-admission")

	InOpenShiftSDNContext(func() {
		g.It("blocks manual creation of Endpoints pointing to the cluster or service network", func() {
			TestEndpointAdmission(g.GinkgoT(), oc)
		})
	})
})

func testOne(t g.GinkgoTInterface, oc *exutil.CLI, client kubernetes.Interface, addrType string, success bool) *corev1.Endpoints {
	networkConfig, err := oc.AdminConfigClient().ConfigV1().Networks().Get(context.Background(), "cluster", metav1.GetOptions{})
	o.Expect(err).NotTo(o.HaveOccurred())

	_, serviceCIDR, err := net.ParseCIDR(networkConfig.Status.ServiceNetwork[0])
	o.Expect(err).NotTo(o.HaveOccurred())
	serviceIP, err := cidr.Host(serviceCIDR, 3)
	o.Expect(err).NotTo(o.HaveOccurred())

	_, clusterCIDR, err := net.ParseCIDR(networkConfig.Status.ClusterNetwork[0].CIDR)
	o.Expect(err).NotTo(o.HaveOccurred())
	clusterIP, err := cidr.Host(clusterCIDR, 3)
	o.Expect(err).NotTo(o.HaveOccurred())

	var exampleAddresses = map[string]string{
		"cluster":  clusterIP.String(),
		"service":  serviceIP.String(),
		"external": "1.2.3.4",
	}

	testEndpoint := &corev1.Endpoints{}
	testEndpoint.GenerateName = "test"
	testEndpoint.Subsets = []corev1.EndpointSubset{
		{
			Addresses: []corev1.EndpointAddress{
				{
					IP: exampleAddresses[addrType],
				},
			},
			Ports: []corev1.EndpointPort{
				{
					Port:     9999,
					Protocol: corev1.ProtocolTCP,
				},
			},
		},
	}

	ep, err := client.CoreV1().Endpoints(oc.Namespace()).Create(context.Background(), testEndpoint, metav1.CreateOptions{})
	if err != nil && success {
		t.Fatalf("unexpected error creating %s network endpoint: %v", addrType, err)
	} else if err == nil && !success {
		t.Fatalf("unexpected success creating %s network endpoint", addrType)
	}
	return ep
}

func TestEndpointAdmission(t g.GinkgoTInterface, oc *exutil.CLI) {
	clusterAdminKubeClient := oc.AdminKubeClient()

	// Cluster admin
	testOne(t, oc, clusterAdminKubeClient, "cluster", true)
	testOne(t, oc, clusterAdminKubeClient, "service", true)
	testOne(t, oc, clusterAdminKubeClient, "external", true)

	// Endpoint controller service account
	serviceAccountClient, _, err := getClientForServiceAccount(clusterAdminKubeClient, rest.AnonymousClientConfig(oc.AdminConfig()), "kube-system", "endpoint-controller")
	if err != nil {
		t.Fatalf("error getting endpoint controller service account: %v", err)
	}
	testOne(t, oc, serviceAccountClient, "cluster", true)
	testOne(t, oc, serviceAccountClient, "service", true)
	testOne(t, oc, serviceAccountClient, "external", true)

	// Project admin
	projectAdminClient := oc.KubeClient()

	testOne(t, oc, projectAdminClient, "cluster", false)
	testOne(t, oc, projectAdminClient, "service", false)
	testOne(t, oc, projectAdminClient, "external", true)

	// User without restricted endpoint permission can't modify IPs but can still do other modifications
	ep := testOne(t, oc, clusterAdminKubeClient, "cluster", true)
	ep.Annotations = map[string]string{"foo": "bar"}
	ep, err = projectAdminClient.CoreV1().Endpoints(oc.Namespace()).Update(context.Background(), ep, metav1.UpdateOptions{})
	if err != nil {
		t.Fatalf("unexpected error updating endpoint annotation: %v", err)
	}

	ep.Subsets[0].Addresses[0].IP = "172.30.0.2"
	ep, err = projectAdminClient.CoreV1().Endpoints(oc.Namespace()).Update(context.Background(), ep, metav1.UpdateOptions{})
	if err == nil {
		t.Fatalf("unexpected success modifying endpoint")
	}
}

func getClientForServiceAccount(adminClient kubernetes.Interface, clientConfig *rest.Config, namespace, name string) (*kubernetes.Clientset, *rest.Config, error) {
	_, err := adminClient.CoreV1().Namespaces().Create(context.Background(), &corev1.Namespace{ObjectMeta: metav1.ObjectMeta{Name: namespace}}, metav1.CreateOptions{})
	if err != nil && !errors.IsAlreadyExists(err) {
		return nil, nil, err
	}

	sa, err := adminClient.CoreV1().ServiceAccounts(namespace).Create(context.Background(), &corev1.ServiceAccount{ObjectMeta: metav1.ObjectMeta{Name: name}}, metav1.CreateOptions{})
	if errors.IsAlreadyExists(err) {
		sa, err = adminClient.CoreV1().ServiceAccounts(namespace).Get(context.Background(), name, metav1.GetOptions{})
	}
	if err != nil {
		return nil, nil, err
	}

	token := ""
	err = wait.Poll(time.Second, 30*time.Second, func() (bool, error) {
		selector := fields.OneTermEqualSelector("type", string(corev1.SecretTypeServiceAccountToken))
		secrets, err := adminClient.CoreV1().Secrets(namespace).List(context.Background(), metav1.ListOptions{FieldSelector: selector.String()})
		if err != nil {
			return false, err
		}
		for _, secret := range secrets.Items {
			if serviceaccount.IsServiceAccountToken(&secret, sa) {
				token = string(secret.Data[corev1.ServiceAccountTokenKey])
				return true, nil
			}
		}
		return false, nil
	})
	if err != nil {
		return nil, nil, err
	}

	saClientConfig := rest.AnonymousClientConfig(clientConfig)
	saClientConfig.BearerToken = token

	kubeClientset, err := kubernetes.NewForConfig(saClientConfig)
	if err != nil {
		return nil, nil, err
	}

	return kubeClientset, saClientConfig, nil
}
