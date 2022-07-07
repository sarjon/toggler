package operator

import "testing"

func TestAnd_AppliesTo(t *testing.T) {
	tests := []struct {
		Operators   []Operator
		Expectation bool
	}{
		{
			Operators: []Operator{
				newMock(true),
				newMock(true),
			},
			Expectation: true,
		},
		{
			Operators: []Operator{
				newMock(true),
				newMock(false),
			},
			Expectation: false,
		},
		{
			Operators: []Operator{
				newMock(false),
				newMock(false),
			},
			Expectation: false,
		},
	}

	for _, test := range tests {
		actual := NewAnd(test.Operators...).AppliesTo("anything")

		if test.Expectation != actual {
			t.Errorf("And(%t).AppliesTo = %t, but expected %t", test.Operators, actual, test.Expectation)
		}
	}
}
