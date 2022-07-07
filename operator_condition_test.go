package toggler

import (
	"context"
	"github.com/sarjon/toggler/operator"
	"testing"
)

func TestOperator_HoldsFor(t *testing.T) {
	tests := []struct {
		Operator operator.Operator
		Context  context.Context
		Holds    bool
	}{
		{
			Operator: newOperatorMock(),
			Context:  context.Background(),
			Holds:    false,
		},
		{
			Operator: newOperatorMock(),
			Context:  context.WithValue(context.Background(), "ip", true),
			Holds:    true,
		},
		{
			Operator: newOperatorMock(),
			Context:  context.WithValue(context.Background(), "ip", false),
			Holds:    false,
		},
	}

	for _, test := range tests {
		actual := NewOperatorCondition("ip", newOperatorMock()).HoldsFor(test.Context)

		if actual != test.Holds {
			t.Errorf("OperatorCondition.HoldsFor = %t, but expected %t", actual, test.Holds)
		}
	}
}

type operatorMock func(value interface{}) bool

func newOperatorMock() operatorMock {
	return func(value interface{}) bool {
		val, ok := value.(bool)
		if ok {
			return val
		}

		return false
	}
}

func (o operatorMock) AppliesTo(value interface{}) bool {
	return o(value)
}
