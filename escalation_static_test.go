package klaxon

import (
	"slices"
	"testing"
	"time"
)

func TestStaticEscalationStrategy(t *testing.T) {
	var (
		disasters = NewStaticEscalationStrategy(Disaster)
		failures  = NewStaticEscalationStrategy(Failure)
		anomalies = NewStaticEscalationStrategy(Anomaly)
	)
	for x := range 100 {
		input := slices.Repeat([]time.Time{time.Now()}, x)
		shouldEqual(t, disasters.CalculateSeverity(input), Disaster)
		shouldEqual(t, failures.CalculateSeverity(input), Failure)
		shouldEqual(t, anomalies.CalculateSeverity(input), Anomaly)
	}
}
