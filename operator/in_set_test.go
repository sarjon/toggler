package operator

import "testing"

func TestInSet_AppliesTo(t *testing.T) {
	tests := []struct {
		Values      []interface{}
		Value       interface{}
		Expectation bool
	}{
		{
			Values:      []interface{}{"A", 1, "B", 2, "C", 3},
			Value:       "B",
			Expectation: true,
		},
		{
			Values:      []interface{}{"A", 1, "B", 2, "C", 3},
			Value:       2,
			Expectation: true,
		},
		{
			Values:      []interface{}{"A", 1, "B", 2, "C", 3},
			Value:       "D",
			Expectation: false,
		},
		{
			Values:      []interface{}{"A", 1, "B", 2, "C", 3},
			Value:       4,
			Expectation: false,
		},
	}

	for _, test := range tests {
		actual := NewInSet(test.Values).AppliesTo(test.Value)

		if actual != test.Expectation {
			t.Errorf("InSet(%v).AppliesTo(%v) = %t, but expected %t", test.Values, test.Value, actual, test.Expectation)
		}
	}
}
