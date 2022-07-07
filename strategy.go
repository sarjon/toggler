package toggler

import "context"

const (
	StrategyAffirmative string = "STRATEGY_AFFIRMATIVE"
	StrategyMajority    string = "STRATEGY_MAJORITY"
	StrategyUnanimous   string = "STRATEGY_UNANIMOUS"
)

type strategy interface {
	execute(ctx context.Context, conditions []Condition) bool
}
