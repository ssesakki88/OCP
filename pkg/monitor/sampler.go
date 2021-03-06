package monitor

import (
	"context"
	"sync"
	"time"

	"github.com/openshift/origin/pkg/monitor/monitorapi"
)

type ConditionalSampler interface {
	ConditionWhenFailing(context.Context, *monitorapi.Condition) SamplerFunc
	WhenFailing(context.Context, *monitorapi.Condition)
}

type sampler struct {
	onFailing *monitorapi.Condition
	interval  time.Duration
	recorder  Recorder
	sampleFn  func(previous bool) (*monitorapi.Condition, bool)

	lock      sync.Mutex
	available bool
}

func NewSampler(recorder Recorder, interval time.Duration, sampleFn func(previous bool) (*monitorapi.Condition, bool)) ConditionalSampler {
	s := &sampler{
		available: true,
		recorder:  recorder,
		interval:  interval,
		sampleFn:  sampleFn,
	}
	return s
}

func (s *sampler) run(ctx context.Context) {
	ticker := time.NewTicker(s.interval)
	defer ticker.Stop()
	var lastInterval int = -1
	for {
		success := s.isAvailable()
		// the sampleFn may take a significant period of time to run.  In such a case, we want our start interval
		// for when a failure started to be the time when the request was first made, not the time when the call
		// returned.  Imagine a timeout set on a DNS lookup of 30s: when the GET finally fails and returns, the outage
		// was actually 30s before.
		startTime := time.Now().UTC()
		condition, ok := s.sampleFn(success)
		if condition != nil {
			s.recorder.RecordAt(startTime, *condition)
		}
		if s.onFailing != nil {
			switch {
			case !success && ok:
				if lastInterval != -1 {
					s.recorder.EndInterval(lastInterval, time.Now().UTC())
				}
			case success && !ok:
				lastInterval = s.recorder.StartInterval(startTime, *s.onFailing)
			}
		}
		s.setAvailable(ok)

		select {
		case <-ticker.C:
		case <-ctx.Done():
			return
		}
	}
}

func (s *sampler) isAvailable() bool {
	s.lock.Lock()
	defer s.lock.Unlock()
	return s.available
}
func (s *sampler) setAvailable(b bool) {
	s.lock.Lock()
	defer s.lock.Unlock()
	s.available = b
}

func (s *sampler) ConditionWhenFailing(ctx context.Context, condition *monitorapi.Condition) SamplerFunc {
	go s.run(ctx)
	return func(_ time.Time) []*monitorapi.Condition {
		if s.isAvailable() {
			return nil
		}
		return []*monitorapi.Condition{condition}
	}
}

func (s *sampler) WhenFailing(ctx context.Context, condition *monitorapi.Condition) {
	s.onFailing = condition
	s.run(ctx)
}
