package frequencyCalculator

import (
	"strings"
	"testing"
)

func TestPrintRepeatedCount(t *testing.T) {
	mappedValues := PrintRepeatedCount("Anshu Budhiraja is Anshu Budhiraja")
	if (mappedValues[strings.ToLower("Anshu")] != 2) || (mappedValues[strings.ToLower("is")] != 1) || (mappedValues[strings.ToLower("Budhiraja")] != 2) {

		t.Fatal("Method not printing correct values")
	}
}
