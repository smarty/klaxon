package klaxon

import "time"

type StaticEscalationStrategy struct {
	severity Severity
}

func NewStaticEscalationStrategy(severity Severity) *StaticEscalationStrategy {
	return &StaticEscalationStrategy{severity: severity}
}

func (this *StaticEscalationStrategy) CalculateSeverity(_ []time.Time) Severity {
	return this.severity
}
