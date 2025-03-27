package klaxon

import (
	"testing"
	"time"
)

func TestWeightedDecayForDisasterAfterOneHourOfContinuousEventsEveryMinute(t *testing.T) {
	now := time.Now()
	Now := func() time.Time { return now }
	escalation := NewWeightedDecayEscalationStrategy(Now, time.Minute, 2.03)

	var severityReadings []Severity
	for eventCount := range 100 {
		var events []time.Time
		for minutesAgo := range eventCount {
			events = append(events, now.Add(-time.Minute*time.Duration(minutesAgo)))
		}
		severityReadings = append(severityReadings, escalation.CalculateSeverity(events))
	}

	shouldEqual(t, len(severityReadings), 100)
	shouldEqual(t, severityReadings, concat(
		repeat(Benign, 1),    // If there are no events: Benign severity
		repeat(Anomaly, 7),   // up to first 7 events are given severity of Anomaly
		repeat(Failure, 52),  // up to 60 events are given severity of Failure
		repeat(Disaster, 40), // beyond 60 events is a Disaster (one full hour of events)
	))
}
func TestWeightedDecayForDisasterAfterOneHourOfContinuousEventsEvery5Minutes(t *testing.T) {
	now := time.Now()
	Now := func() time.Time { return now }
	escalation := NewWeightedDecayEscalationStrategy(Now, time.Minute, .25)

	var severityReadings []Severity
	for eventCount := range 20 {
		var events []time.Time
		for multiplier := range eventCount {
			events = append(events, now.Add(-time.Minute*(5*time.Duration(multiplier))))
		}
		severityReadings = append(severityReadings, escalation.CalculateSeverity(events))
	}

	shouldEqual(t, len(severityReadings), 20)
	shouldEqual(t, severityReadings, concat(
		repeat(Benign, 1),   // If there are no events: Benign severity
		repeat(Anomaly, 4),  // up to first 7 events are given severity of Anomaly
		repeat(Failure, 11), // up to 60 events are given severity of Failure
		repeat(Disaster, 4), // beyond 60 events is a Disaster (one full hour of events)
	))
}

func repeat[T any](v T, count int) (result []T) {
	for range count {
		result = append(result, v)
	}
	return result
}
func concat[T any](slices ...[]T) (result []T) {
	for _, s := range slices {
		result = append(result, s...)
	}
	return result
}
