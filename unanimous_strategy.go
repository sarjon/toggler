package toggler

import (
	"context"
)

// unanimousStrategy strategy checks that all conditions hold.
type unanimousStrategy struct {
}

func newUnanimousStrategy() *unanimousStrategy {
	return &unanimousStrategy{}
}

func (u *unanimousStrategy) execute(ctx context.Context, conditions []Condition) bool {
	for _, condition := range conditions {
		if !condition.HoldsFor(ctx) {
			return false
		}
	}

	return true
}
