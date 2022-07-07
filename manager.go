package toggler

import (
	"context"
	"errors"
)

var (
	ErrToggleNotFound               = errors.New("toggle not found")
	ErrToggleNotConfigured          = errors.New("toggle not configured")
	ErrToggleActivationNotSupported = errors.New("toggle activation not supported")
	ErrToggleStrategyNotSupported   = errors.New("toggle strategy not supported")
)

type Manager struct {
	storage    Storage
	strategies map[string]strategy
}

type ManagerArgs struct {
	Storage Storage
}

func NewManager(args *ManagerArgs) *Manager {
	return &Manager{
		storage: args.Storage,
		strategies: map[string]strategy{
			StrategyAffirmative: newAffirmativeStrategy(),
			StrategyMajority:    newMajorityStrategy(),
			StrategyUnanimous:   newUnanimousStrategy(),
		},
	}
}

func (m *Manager) Active(name string, ctx context.Context) (bool, error) {
	toggle, found := m.storage.Get(name)
	if !found {
		return false, ErrToggleNotFound
	}

	if toggle.IsZero() {
		return false, ErrToggleNotConfigured
	}

	switch toggle.Activation() {
	case ActiveAlways:
		return true, nil
	case ActiveNever:
		return false, nil
	case ActiveConditionally:
		strategy, found := m.strategies[toggle.Strategy()]
		if !found {
			return false, ErrToggleStrategyNotSupported
		}

		return strategy.execute(ctx, toggle.Conditions()), nil
	}

	return false, ErrToggleActivationNotSupported
}
