package engine

import (
	"testing"

	"github.com/phsm/pandora/lib/ginkgoutil"
	"github.com/phsm/pandora/lib/monitoring"
)

func TestEngine(t *testing.T) {
	ginkgoutil.RunSuite(t, "Engine Suite")
}

func newTestMetrics() Metrics {
	return Metrics{
		&monitoring.Counter{},
		&monitoring.Counter{},
		&monitoring.Counter{},
		&monitoring.Counter{},
	}
}
