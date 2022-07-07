package operator

// And operator evaluates whether all operators apply to a given value.
type And struct {
	operators []Operator
}

func NewAnd(operators ...Operator) *And {
	return &And{
		operators: operators,
	}
}

func (a *And) AppliesTo(value interface{}) bool {
	for _, operator := range a.operators {
		if !operator.AppliesTo(value) {
			return false
		}
	}

	return true
}
