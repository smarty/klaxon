package klaxon

import "time"

type Severity int

const (
	Benign   Severity = 0
	Anomaly  Severity = 1
	Failure  Severity = 2
	Disaster Severity = 3
)

type Monitor interface {
	Monitor(Severity)
}

type EscalationStrategy interface {
	CalculateSeverity(events []time.Time) Severity
}

type Sensor interface {
	Reset()
	Record(event time.Time) Severity
}
