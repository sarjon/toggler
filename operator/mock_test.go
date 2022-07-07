package operator

import "testing"

func TestMock_AppliesTo(t *testing.T) {
	tests := []struct {
		Value       bool
		Expectation bool
	}{
		{
			Value:       false,
			Expectation: false,
		},
		{
			Value:       true,
			Expectation: true,
		},
	}

	for _, test := range tests {
		actual := newMock(test.Value).AppliesTo("anything")

		if test.Expectation != actual {
			t.Errorf("mock(%t).AppliesTo = %t, but expected %t", test.Value, actual, test.Expectation)
		}
	}
}
