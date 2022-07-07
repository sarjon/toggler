package toggler

import "context"

type Condition interface {
	HoldsFor(ctx context.Context) bool
}

type mockCondition struct {
	holds bool
}

func newMockCondition(holds bool) *mockCondition {
	return &mockCondition{
		holds: holds,
	}
}

func (m *mockCondition) HoldsFor(ctx context.Context) bool {
	return m.holds
}
