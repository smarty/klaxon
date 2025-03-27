package klaxon

import "time"

type sensor struct {
	strategy EscalationStrategy
	monitor  Monitor
	history  []time.Time
}

func NewSensor(strategy EscalationStrategy, monitor Monitor) Sensor {
	return &sensor{strategy: strategy, monitor: monitor}
}
func (this *sensor) Reset() {
	if len(this.history) > 0 {
		this.history = this.history[:0]
	}
}
func (this *sensor) Record(event time.Time) (result Severity) {
	if len(this.history) == maxHistoryCount {
		this.history = this.history[1:]
	}
	this.history = append(this.history, event)
	result = this.strategy.CalculateSeverity(this.history)
	this.monitor.Monitor(result)
	return result
}

const maxHistoryCount = 1024
