package klaxon

import "time"

// WeightedDecayEscalationStrategy evaluates the severity of a series of events by calculating a 'weight' score for each
// based on recency, as well as a frequency scale.
type WeightedDecayEscalationStrategy struct {
	clock        func() time.Time
	ageUnit      time.Duration // Evaluate an event's age in terms of this unit (time.Millisecond, time.Second, etc...).
	decayDivisor float64       // The larger the value, the less weight older events will be assessed.
}

func NewWeightedDecayEscalationStrategy(
	clock func() time.Time, ageUnit time.Duration, decayDivisor float64,
) *WeightedDecayEscalationStrategy {
	return &WeightedDecayEscalationStrategy{
		clock:        clock,
		ageUnit:      ageUnit,
		decayDivisor: decayDivisor,
	}
}
func (this *WeightedDecayEscalationStrategy) CalculateSeverity(events []time.Time) Severity {
	return Severity(this.calculateCombinedWeight(events))
}
func (this *WeightedDecayEscalationStrategy) calculateCombinedWeight(events []time.Time) (result float64) {
	now := this.clock()
	for _, event := range events {
		age := in(this.ageUnit, now.Sub(event))
		weight := 1.0 / (1.0 + this.decayDivisor*age)
		result += weight
		if Severity(result) >= Disaster {
			break
		}
	}
	return result
}

// in returns d in terms of the provided unit.
// in(time.Second, d) is equivalent to d.Seconds().
// in(time.Minute, d) is equivalent to d.Minutes().
// in(time.Hour, d) is equivalent to d.Hours().
func in(unit, d time.Duration) float64 {
	return float64(d/unit) + float64(d%unit)/float64(unit)
}
