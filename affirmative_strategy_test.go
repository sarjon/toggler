package toggler

import (
	"context"
	"testing"
)

func TestAffirmativeStrategy_holdsFor(t *testing.T) {
	tests := []struct {
		Conditions  []Condition
		Expectation bool
	}{
		{
			Conditions: []Condition{
				newMockCondition(false),
				newMockCondition(true),
				newMockCondition(false),
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
			},
			Expectation: true,
		},
	}

	for _, test := range tests {
		actual := newAffirmativeStrategy().execute(context.Background(), test.Conditions)

		if actual != test.Expectation {
			t.Errorf("affirmativeStrategy returned %t, but %t expected.", actual, test.Expectation)
		}
	}
}
