package toggler

import (
	"context"
	"github.com/sarjon/toggler/operator"
)

type OperatorCondition struct {
	operator operator.Operator
	key      string
}

func NewOperatorCondition(key string, operator operator.Operator) *OperatorCondition {
	return &OperatorCondition{
		operator: operator,
		key:      key,
	}
}

func (o *OperatorCondition) HoldsFor(ctx context.Context) bool {
	value := ctx.Value(o.key)
	if value == nil {
		return false
	}

	return o.operator.AppliesTo(value)
}
