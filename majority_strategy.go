package toggler

import (
	"context"
)

// majorityStrategy strategy checks if the majority of conditions hold.
type majorityStrategy struct {
}

func newMajorityStrategy() *majorityStrategy {
	return &majorityStrategy{}
}

func (m *majorityStrategy) execute(ctx context.Context, conditions []Condition) bool {
	positives, negatives := 0, 0

	for _, condition := range conditions {
		if condition.HoldsFor(ctx) {
			positives++
			continue
		}

		negatives++
	}

	return positives > negatives
}
