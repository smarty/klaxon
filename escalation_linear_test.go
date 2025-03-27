package klaxon

import (
	"testing"
	"time"
)

func TestNoEvents_Benign(t *testing.T) {
	start := time.Now()
	Now := func() time.Time { return start }

	escalation := NewLinearEscalationStrategy(Now, time.Nanosecond, 1)
	severity := escalation.CalculateSeverity([]time.Time(nil))
	shouldEqual(t, severity, Benign)
}
func TestSingleEvent_Anomaly(t *testing.T) {
	start := time.Now()
	Now := func() time.Time { return start }

	escalation := NewLinearEscalationStrategy(Now, time.Nanosecond, 1)
	severity := escalation.CalculateSeverity([]time.Time{start})
	shouldEqual(t, severity, Anomaly)
}
func TestTwoEvents_Failure(t *testing.T) {
	start := time.Now()
	Now := func() time.Time { return start }

	escalation := NewLinearEscalationStrategy(Now, time.Second, 1)
	severity := escalation.CalculateSeverity([]time.Time{start, start})
	shouldEqual(t, severity, Failure)
}
func TestThreeEvents_Disaster(t *testing.T) {
	start := time.Now()
	Now := func() time.Time { return start }

	escalation := NewLinearEscalationStrategy(Now, time.Nanosecond, 1)
	severity := escalation.CalculateSeverity([]time.Time{start, start, start})
	shouldEqual(t, severity, Disaster)
}
func TestMoreThanThreeEvents_Disaster(t *testing.T) {
	start := time.Now()
	Now := func() time.Time { return start }

	escalation := NewLinearEscalationStrategy(Now, time.Nanosecond, 1)
	severity := escalation.CalculateSeverity([]time.Time{start, start, start, start})
	shouldEqual(t, severity, Disaster)
}
func TestPastDisaster(t *testing.T) {
	start := time.Now()
	Now := func() time.Time { return start }

	escalation := NewLinearEscalationStrategy(Now, time.Hour, 1)
	severity := escalation.CalculateSeverity([]time.Time{
		start.Add(time.Minute * 0),
		start.Add(time.Minute * 1),
		start.Add(time.Minute * 2), // three in a row, past disaster

		start.Add(time.Hour * 1).Add(time.Minute * 0),
		start.Add(time.Hour * 1).Add(time.Minute * 1), // two in a row, more recent failure
	})
	shouldEqual(t, severity, Disaster)
}
