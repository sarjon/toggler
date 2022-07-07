package operator

import "testing"

func TestRegex_AppliesTo(t *testing.T) {
	tests := []struct {
		Regex       string
		Value       interface{}
		Expectation bool
	}{
		{
			Regex:       "^begin",
			Value:       "begins with this",
			Expectation: true,
		},
		{
			Regex:       "\\d",
			Value:       "5",
			Expectation: true,
		},
		{
			Regex:       "\\d",
			Value:       "A",
			Expectation: false,
		},
	}

	for _, test := range tests {
		actual := NewRegex(test.Regex).AppliesTo(test.Value)

		if test.Expectation != actual {
			t.Errorf("Regex(%s).AppliesTo(%v) = %t, but expected %t", test.Regex, test.Value, actual, test.Expectation)
		}
	}
}
