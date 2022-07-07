package toggler

import (
	"context"
	"testing"
)

func TestMajorityStrategy_holdsFor(t *testing.T) {
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
			Expectation: false,
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
			},
			Expectation: false,
		},
		{
			Conditions: []Condition{
				newMockCondition(true),
				newMockCondition(false),
				newMockCondition(true),
			},
			Expectation: true,
		},
	}

	for _, test := range tests {
		actual := newMajorityStrategy().execute(context.Background(), test.Conditions)

		if actual != test.Expectation {
			t.Errorf("majorityStrategy returned %t, but %t expected.", actual, test.Expectation)
		}
	}
}
