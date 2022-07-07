package toggler

import (
	"context"
	"testing"
)

func TestUnanimousStrategy_holdsFor(t *testing.T) {
	tests := []struct {
		Conditions  []Condition
		Expectation bool
	}{
		{
			Conditions: []Condition{
				newMockCondition(true),
				newMockCondition(true),
				newMockCondition(true),
			},
			Expectation: true,
		},
		{
			Conditions: []Condition{
				newMockCondition(false),
				newMockCondition(false),
			},
			Expectation: false,
		},
		{
			Conditions: []Condition{
				newMockCondition(true),
				newMockCondition(false),
				newMockCondition(true),
			},
			Expectation: false,
		},
	}

	for _, test := range tests {
		actual := newUnanimousStrategy().execute(context.Background(), test.Conditions)

		if actual != test.Expectation {
			t.Errorf("unanimousStrategy returned %t, but %t expected.", actual, test.Expectation)
		}
	}
}
