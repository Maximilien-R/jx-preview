package tracing

import (
	"time"

	"github.com/jenkins-x/jx-logging/v3/pkg/log"
)

// TimeIt logs the start and end time of a named operation.
// Usage: defer TimeIt("my operation")()
func TimeIt(name string) func() {
	start := time.Now()
	return func() {
		end := time.Now()
		log.Logger().Infof(
			"[TIMEIT] %s took %v",
			name,
			end.Sub(start),
		)
	}
}
