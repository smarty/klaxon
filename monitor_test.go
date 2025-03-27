package klaxon

import "testing"

func TestMetricsMonitorFixture(t *testing.T) {
	var (
		anomalies = NewFakeCounter()
		failures  = NewFakeCounter()
		disasters = NewFakeCounter()
		monitor   = NewMetricsMonitor(anomalies, failures, disasters)
	)

	monitor.Monitor(Benign) // no-op

	monitor.Monitor(Anomaly)
	monitor.Monitor(Anomaly)

	monitor.Monitor(Failure)
	monitor.Monitor(Failure)
	monitor.Monitor(Failure)

	monitor.Monitor(Disaster)
	monitor.Monitor(Disaster)
	monitor.Monitor(Disaster)
	monitor.Monitor(Disaster)

	shouldEqual(t, anomalies.value, 2)
	shouldEqual(t, failures.value, 3)
	shouldEqual(t, disasters.value, 4)
}
