package toggler

import (
	"context"
)

// affirmativeStrategy strategy checks if at least one condition holds.
type affirmativeStrategy struct {
}

func newAffirmativeStrategy() *affirmativeStrategy {
	return &affirmativeStrategy{}
}

func (a *affirmativeStrategy) execute(ctx context.Context, conditions []Condition) bool {
	for _, c := range conditions {
		if c.HoldsFor(ctx) {
			return true
		}
	}

	return false
}
