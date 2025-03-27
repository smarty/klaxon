package klaxon

import (
	"testing"
	"time"
)

func TestSensor(t *testing.T) {
	monitor := NewFakeMonitor()
	instrument := NewSensor(NewFakeEscalationStrategy(), monitor)

	severity := instrument.Record(time.Now())

	shouldEqual(t, severity, Anomaly)
	shouldEqual(t, monitor.monitored, []Severity{Anomaly})

	severity = instrument.Record(time.Now())

	shouldEqual(t, severity, Failure)
	shouldEqual(t, monitor.monitored, []Severity{Anomaly, Failure})

	severity = instrument.Record(time.Now())

	shouldEqual(t, severity, Disaster)
	shouldEqual(t, monitor.monitored, []Severity{Anomaly, Failure, Disaster})

	severity = instrument.Record(time.Now())

	shouldEqual(t, severity, Disaster)
	shouldEqual(t, monitor.monitored, []Severity{Anomaly, Failure, Disaster, Disaster})

	instrument.Reset()

	severity = instrument.Record(time.Now())

	shouldEqual(t, severity, Anomaly)
	shouldEqual(t, monitor.monitored, []Severity{Anomaly, Failure, Disaster, Disaster, Anomaly})
}

func TestSensorEventsCapped(t *testing.T) {
	strategy := NewFakeEscalationStrategy()
	instrument := NewSensor(strategy, NewFakeMonitor())
	for range maxHistoryCount {
		_ = instrument.Record(time.Now())
	}
	shouldEqual(t, len(strategy.events), maxHistoryCount)
	_ = instrument.Record(time.Now())
	shouldEqual(t, len(strategy.events), maxHistoryCount)
}
