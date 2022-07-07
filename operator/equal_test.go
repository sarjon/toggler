package operator

import "testing"

func TestEqual_AppliesTo(t *testing.T) {
	tests := []struct {
		ValueA      interface{}
		ValueB      interface{}
		Expectation bool
	}{
		{
			ValueA:      3,
			ValueB:      3,
			Expectation: true,
		},
		{
			ValueA:      "ABC",
			ValueB:      "ABC",
			Expectation: true,
		},
		{
			ValueA:      123,
			ValueB:      321,
			Expectation: false,
		},
	}

	for _, test := range tests {
		operator := NewEqual(test.ValueA)

		actual := operator.AppliesTo(test.ValueB)

		if actual != test.Expectation {
			t.Errorf("equal(%s).AppliesTo(%s) = %t, but expected %t", test.ValueA, test.ValueB, actual, test.Expectation)
		}
	}
}
