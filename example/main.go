package main

import (
	"log"
	"time"

	"github.com/smarty/klaxon/v2"
)

func main() {
	log.SetFlags(0)
	log.SetPrefix(">>> ")

	var (
		anomalies = NewCounter()
		failures  = NewCounter()
		disasters = NewCounter()
	)

	now := time.Now()
	clock := func() time.Time { return now }

	monitor := klaxon.NewMetricsMonitor(anomalies, failures, disasters)
	strategy := klaxon.NewWeightedDecayEscalationStrategy(clock, time.Minute, .25)
	sensor := klaxon.NewSensor(strategy, monitor)

	for range 25 {
		log.Println(now.Format(time.DateTime), "-- Severity:", sensor.Record(now))
		now = now.Add(time.Minute * 5)
	}

	log.Println("Anomalies:", anomalies.Value)
	log.Println("Failures:", failures.Value)
	log.Println("Disasters:", disasters.Value)
}

type Counter struct{ Value int }

func NewCounter() *Counter       { return &Counter{} }
func (this *Counter) Increment() { this.Value++ }
