package klaxon

type metricsMonitor struct {
	anomalies counter
	failures  counter
	disasters counter
}

func NewMetricsMonitor(anomalies, failures, disasters counter) Monitor {
	return &metricsMonitor{
		anomalies: anomalies,
		failures:  failures,
		disasters: disasters,
	}
}

func (this *metricsMonitor) Monitor(severity Severity) {
	switch severity {
	case Benign:
		// no-op
	case Anomaly:
		this.anomalies.Increment()
	case Failure:
		this.failures.Increment()
	case Disaster:
		this.disasters.Increment()
	}
}

type counter interface {
	Increment()
}
