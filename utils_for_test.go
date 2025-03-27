package klaxon

import (
	"reflect"
	"testing"
	"time"
)

func shouldEqual(t *testing.T, actual, expected any) {
	if reflect.DeepEqual(actual, expected) {
		return
	}
	t.Helper()
	t.Errorf("%v != %v", actual, expected)
}

//////////////////////////////////////////////////////

type FakeCounter struct{ value int }

func NewFakeCounter() *FakeCounter {
	return &FakeCounter{}
}

func (this *FakeCounter) Increment() { this.value++ }

//////////////////////////////////////////////////////

type FakeEscalationStrategy struct{}

func NewFakeEscalationStrategy() *FakeEscalationStrategy {
	return &FakeEscalationStrategy{}
}

func (this *FakeEscalationStrategy) CalculateSeverity(events []time.Time) Severity {
	return min(Disaster, Severity(len(events)))
}

//////////////////////////////////////////////////////

type FakeMonitor struct {
	monitored []Severity
}

func NewFakeMonitor() *FakeMonitor {
	return &FakeMonitor{}
}

func (this *FakeMonitor) Monitor(severity Severity) {
	this.monitored = append(this.monitored, severity)
}

//////////////////////////////////////////////////////
