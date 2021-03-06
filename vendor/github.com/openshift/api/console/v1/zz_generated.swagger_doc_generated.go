package v1

// This file contains a collection of methods that can be used from go-restful to
// generate Swagger API documentation for its models. Please read this PR for more
// information on the implementation: https://github.com/emicklei/go-restful/pull/215
//
// TODOs are ignored from the parser (e.g. TODO(andronat):... || TODO:...) if and only if
// they are on one line! For multiple line or blocks that you want to ignore use ---.
// Any context after a --- is ignored.
//
// Those methods can be generated by using hack/update-swagger-docs.sh

// AUTO-GENERATED FUNCTIONS START HERE
var map_Link = map[string]string{
	"":     "Represents a standard link that could be generated in HTML",
	"text": "text is the display text for the link",
	"href": "href is the absolute secure URL for the link (must use https)",
}

func (Link) SwaggerDoc() map[string]string {
	return map_Link
}

var map_CLIDownloadLink = map[string]string{
	"text": "text is the display text for the link",
	"href": "href is the absolute secure URL for the link (must use https)",
}

func (CLIDownloadLink) SwaggerDoc() map[string]string {
	return map_CLIDownloadLink
}

var map_ConsoleCLIDownload = map[string]string{
	"": "ConsoleCLIDownload is an extension for configuring openshift web console command line interface (CLI) downloads.\n\nCompatibility level 1: Stable within a major release for a minimum of 12 months or 3 minor releases (whichever is longer).",
}

func (ConsoleCLIDownload) SwaggerDoc() map[string]string {
	return map_ConsoleCLIDownload
}

var map_ConsoleCLIDownloadList = map[string]string{
	"": "Compatibility level 1: Stable within a major release for a minimum of 12 months or 3 minor releases (whichever is longer).",
}

func (ConsoleCLIDownloadList) SwaggerDoc() map[string]string {
	return map_ConsoleCLIDownloadList
}

var map_ConsoleCLIDownloadSpec = map[string]string{
	"":            "ConsoleCLIDownloadSpec is the desired cli download configuration.",
	"displayName": "displayName is the display name of the CLI download.",
	"description": "description is the description of the CLI download (can include markdown).",
	"links":       "links is a list of objects that provide CLI download link details.",
}

func (ConsoleCLIDownloadSpec) SwaggerDoc() map[string]string {
	return map_ConsoleCLIDownloadSpec
}

var map_ConsoleExternalLogLink = map[string]string{
	"": "ConsoleExternalLogLink is an extension for customizing OpenShift web console log links.\n\nCompatibility level 1: Stable within a major release for a minimum of 12 months or 3 minor releases (whichever is longer).",
}

func (ConsoleExternalLogLink) SwaggerDoc() map[string]string {
	return map_ConsoleExternalLogLink
}

var map_ConsoleExternalLogLinkList = map[string]string{
	"": "Compatibility level 1: Stable within a major release for a minimum of 12 months or 3 minor releases (whichever is longer).",
}

func (ConsoleExternalLogLinkList) SwaggerDoc() map[string]string {
	return map_ConsoleExternalLogLinkList
}

var map_ConsoleExternalLogLinkSpec = map[string]string{
	"":                "ConsoleExternalLogLinkSpec is the desired log link configuration. The log link will appear on the logs tab of the pod details page.",
	"text":            "text is the display text for the link",
	"hrefTemplate":    "hrefTemplate is an absolute secure URL (must use https) for the log link including variables to be replaced. Variables are specified in the URL with the format ${variableName}, for instance, ${containerName} and will be replaced with the corresponding values from the resource. Resource is a pod. Supported variables are: - ${resourceName} - name of the resource which containes the logs - ${resourceUID} - UID of the resource which contains the logs\n              - e.g. `11111111-2222-3333-4444-555555555555`\n- ${containerName} - name of the resource's container that contains the logs - ${resourceNamespace} - namespace of the resource that contains the logs - ${resourceNamespaceUID} - namespace UID of the resource that contains the logs - ${podLabels} - JSON representation of labels matching the pod with the logs\n            - e.g. `{\"key1\":\"value1\",\"key2\":\"value2\"}`\n\ne.g., https://example.com/logs?resourceName=${resourceName}&containerName=${containerName}&resourceNamespace=${resourceNamespace}&podLabels=${podLabels}",
	"namespaceFilter": "namespaceFilter is a regular expression used to restrict a log link to a matching set of namespaces (e.g., `^openshift-`). The string is converted into a regular expression using the JavaScript RegExp constructor. If not specified, links will be displayed for all the namespaces.",
}

func (ConsoleExternalLogLinkSpec) SwaggerDoc() map[string]string {
	return map_ConsoleExternalLogLinkSpec
}

var map_ApplicationMenuSpec = map[string]string{
	"":         "ApplicationMenuSpec is the specification of the desired section and icon used for the link in the application menu.",
	"section":  "section is the section of the application menu in which the link should appear. This can be any text that will appear as a subheading in the application menu dropdown. A new section will be created if the text does not match text of an existing section.",
	"imageURL": "imageUrl is the URL for the icon used in front of the link in the application menu. The URL must be an HTTPS URL or a Data URI. The image should be square and will be shown at 24x24 pixels.",
}

func (ApplicationMenuSpec) SwaggerDoc() map[string]string {
	return map_ApplicationMenuSpec
}

var map_ConsoleLink = map[string]string{
	"": "ConsoleLink is an extension for customizing OpenShift web console links.\n\nCompatibility level 1: Stable within a major release for a minimum of 12 months or 3 minor releases (whichever is longer).",
}

func (ConsoleLink) SwaggerDoc() map[string]string {
	return map_ConsoleLink
}

var map_ConsoleLinkList = map[string]string{
	"": "Compatibility level 1: Stable within a major release for a minimum of 12 months or 3 minor releases (whichever is longer).",
}

func (ConsoleLinkList) SwaggerDoc() map[string]string {
	return map_ConsoleLinkList
}

var map_ConsoleLinkSpec = map[string]string{
	"":                   "ConsoleLinkSpec is the desired console link configuration.",
	"location":           "location determines which location in the console the link will be appended to (ApplicationMenu, HelpMenu, UserMenu, NamespaceDashboard).",
	"applicationMenu":    "applicationMenu holds information about section and icon used for the link in the application menu, and it is applicable only when location is set to ApplicationMenu.",
	"namespaceDashboard": "namespaceDashboard holds information about namespaces in which the dashboard link should appear, and it is applicable only when location is set to NamespaceDashboard. If not specified, the link will appear in all namespaces.",
}

func (ConsoleLinkSpec) SwaggerDoc() map[string]string {
	return map_ConsoleLinkSpec
}

var map_NamespaceDashboardSpec = map[string]string{
	"":                  "NamespaceDashboardSpec is a specification of namespaces in which the dashboard link should appear. If both namespaces and namespaceSelector are specified, the link will appear in namespaces that match either",
	"namespaces":        "namespaces is an array of namespace names in which the dashboard link should appear.",
	"namespaceSelector": "namespaceSelector is used to select the Namespaces that should contain dashboard link by label. If the namespace labels match, dashboard link will be shown for the namespaces.",
}

func (NamespaceDashboardSpec) SwaggerDoc() map[string]string {
	return map_NamespaceDashboardSpec
}

var map_ConsoleNotification = map[string]string{
	"": "ConsoleNotification is the extension for configuring openshift web console notifications.\n\nCompatibility level 1: Stable within a major release for a minimum of 12 months or 3 minor releases (whichever is longer).",
}

func (ConsoleNotification) SwaggerDoc() map[string]string {
	return map_ConsoleNotification
}

var map_ConsoleNotificationList = map[string]string{
	"": "Compatibility level 1: Stable within a major release for a minimum of 12 months or 3 minor releases (whichever is longer).",
}

func (ConsoleNotificationList) SwaggerDoc() map[string]string {
	return map_ConsoleNotificationList
}

var map_ConsoleNotificationSpec = map[string]string{
	"":                "ConsoleNotificationSpec is the desired console notification configuration.",
	"text":            "text is the visible text of the notification.",
	"location":        "location is the location of the notification in the console. Valid values are: \"BannerTop\", \"BannerBottom\", \"BannerTopBottom\".",
	"link":            "link is an object that holds notification link details.",
	"color":           "color is the color of the text for the notification as CSS data type color.",
	"backgroundColor": "backgroundColor is the color of the background for the notification as CSS data type color.",
}

func (ConsoleNotificationSpec) SwaggerDoc() map[string]string {
	return map_ConsoleNotificationSpec
}

var map_ConsoleQuickStart = map[string]string{
	"": "ConsoleQuickStart is an extension for guiding user through various workflows in the OpenShift web console.\n\nCompatibility level 1: Stable within a major release for a minimum of 12 months or 3 minor releases (whichever is longer).",
}

func (ConsoleQuickStart) SwaggerDoc() map[string]string {
	return map_ConsoleQuickStart
}

var map_ConsoleQuickStartList = map[string]string{
	"": "Compatibility level 1: Stable within a major release for a minimum of 12 months or 3 minor releases (whichever is longer).",
}

func (ConsoleQuickStartList) SwaggerDoc() map[string]string {
	return map_ConsoleQuickStartList
}

var map_ConsoleQuickStartSpec = map[string]string{
	"":                      "ConsoleQuickStartSpec is the desired quick start configuration.",
	"displayName":           "displayName is the display name of the Quick Start.",
	"icon":                  "icon is a base64 encoded image that will be displayed beside the Quick Start display name. The icon should be an vector image for easy scaling. The size of the icon should be 40x40.",
	"tags":                  "tags is a list of strings that describe the Quick Start.",
	"durationMinutes":       "durationMinutes describes approximately how many minutes it will take to complete the Quick Start.",
	"description":           "description is the description of the Quick Start. (includes markdown)",
	"prerequisites":         "prerequisites contains all prerequisites that need to be met before taking a Quick Start. (includes markdown)",
	"introduction":          "introduction describes the purpose of the Quick Start. (includes markdown)",
	"tasks":                 "tasks is the list of steps the user has to perform to complete the Quick Start.",
	"conclusion":            "conclusion sums up the Quick Start and suggests the possible next steps. (includes markdown)",
	"nextQuickStart":        "nextQuickStart is a list of the following Quick Starts, suggested for the user to try.",
	"accessReviewResources": "accessReviewResources contains a list of resources that the user's access will be reviewed against in order for the user to complete the Quick Start. The Quick Start will be hidden if any of the access reviews fail.",
}

func (ConsoleQuickStartSpec) SwaggerDoc() map[string]string {
	return map_ConsoleQuickStartSpec
}

var map_ConsoleQuickStartTask = map[string]string{
	"":            "ConsoleQuickStartTask is a single step in a Quick Start.",
	"title":       "title describes the task and is displayed as a step heading.",
	"description": "description describes the steps needed to complete the task. (includes markdown)",
	"review":      "review contains instructions to validate the task is complete. The user will select 'Yes' or 'No'. using a radio button, which indicates whether the step was completed successfully.",
	"summary":     "summary contains information about the passed step.",
}

func (ConsoleQuickStartTask) SwaggerDoc() map[string]string {
	return map_ConsoleQuickStartTask
}

var map_ConsoleQuickStartTaskReview = map[string]string{
	"":               "ConsoleQuickStartTaskReview contains instructions that validate a task was completed successfully.",
	"instructions":   "instructions contains steps that user needs to take in order to validate his work after going through a task. (includes markdown)",
	"failedTaskHelp": "failedTaskHelp contains suggestions for a failed task review and is shown at the end of task. (includes markdown)",
}

func (ConsoleQuickStartTaskReview) SwaggerDoc() map[string]string {
	return map_ConsoleQuickStartTaskReview
}

var map_ConsoleQuickStartTaskSummary = map[string]string{
	"":        "ConsoleQuickStartTaskSummary contains information about a passed step.",
	"success": "success describes the succesfully passed task.",
	"failed":  "failed briefly describes the unsuccessfully passed task. (includes markdown)",
}

func (ConsoleQuickStartTaskSummary) SwaggerDoc() map[string]string {
	return map_ConsoleQuickStartTaskSummary
}

var map_ConsoleYAMLSample = map[string]string{
	"": "ConsoleYAMLSample is an extension for customizing OpenShift web console YAML samples.\n\nCompatibility level 1: Stable within a major release for a minimum of 12 months or 3 minor releases (whichever is longer).",
}

func (ConsoleYAMLSample) SwaggerDoc() map[string]string {
	return map_ConsoleYAMLSample
}

var map_ConsoleYAMLSampleList = map[string]string{
	"": "Compatibility level 1: Stable within a major release for a minimum of 12 months or 3 minor releases (whichever is longer).",
}

func (ConsoleYAMLSampleList) SwaggerDoc() map[string]string {
	return map_ConsoleYAMLSampleList
}

var map_ConsoleYAMLSampleSpec = map[string]string{
	"":               "ConsoleYAMLSampleSpec is the desired YAML sample configuration. Samples will appear with their descriptions in a samples sidebar when creating a resources in the web console.",
	"targetResource": "targetResource contains apiVersion and kind of the resource YAML sample is representating.",
	"title":          "title of the YAML sample.",
	"description":    "description of the YAML sample.",
	"yaml":           "yaml is the YAML sample to display.",
	"snippet":        "snippet indicates that the YAML sample is not the full YAML resource definition, but a fragment that can be inserted into the existing YAML document at the user's cursor.",
}

func (ConsoleYAMLSampleSpec) SwaggerDoc() map[string]string {
	return map_ConsoleYAMLSampleSpec
}

// AUTO-GENERATED FUNCTIONS END HERE
