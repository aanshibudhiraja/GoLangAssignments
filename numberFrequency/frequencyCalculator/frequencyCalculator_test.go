package frequencyCalculator

import "testing"

// TestPrintNumberFrequency method
func TestPrintNumberFrequency(t *testing.T) {
	tr := PrintNumberfrequency(1234556121212129871)
	if !tr {
		t.Fatal("Unable to count numbers")
	}
}
