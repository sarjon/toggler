package operator

import (
	"testing"
)

func TestPercentage_AppliesTo(t *testing.T) {
	tests := []struct {
		Percent     int
		Value       interface{}
		Expectation bool
	}{
		{
			Percent:     35,
			Value:       "sarjon@example.com",
			Expectation: true,
		},
		{
			Percent:     100,
			Value:       1,
			Expectation: false,
		},
	}

	for _, test := range tests {
		percentage := NewPercentage(test.Percent)
		actual := percentage.AppliesTo(test.Value)

		if actual != test.Expectation {
			t.Errorf("Percentage(%d).AppliesTo(%s) = %t, but expected %t", test.Percent, test.Value, actual, test.Expectation)
		}
	}

}
