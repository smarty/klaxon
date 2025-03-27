package klaxon

import "time"

type LinearEscalationStrategy struct {
	clock          func() time.Time
	duration       time.Duration
	frequencyScale int
}

func NewLinearEscalationStrategy(clock func() time.Time, duration time.Duration, frequencyScale int) *LinearEscalationStrategy {
	return &LinearEscalationStrategy{
		clock:          clock,
		duration:       duration,
		frequencyScale: frequencyScale,
	}
}
func (this *LinearEscalationStrategy) CalculateSeverity(events []time.Time) (result Severity) {
	windowStart := this.clock().Add(-this.duration)
	for i := len(events) - 1; i >= 0; i-- {
		if events[i].After(windowStart) {
			result++
		}
		if result >= Disaster {
			break
		}
	}
	return result
}
