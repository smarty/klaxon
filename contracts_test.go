package klaxon

import "testing"

func TestSeverity_String(t *testing.T) {
	shouldEqual(t, "Benign", Benign.String())
	shouldEqual(t, "Anomaly", Anomaly.String())
	shouldEqual(t, "Failure", Failure.String())
	shouldEqual(t, "Disaster", Disaster.String())
	shouldEqual(t, "Unknown", Severity(42).String())
}
