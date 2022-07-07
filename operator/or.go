package operator

// Or operator evaluates if at least one operator applies to given value.
type Or struct {
	operators []Operator
}

func NewOr(operators ...Operator) *Or {
	return &Or{
		operators: operators,
	}
}

func (o *Or) AppliesTo(value interface{}) bool {
	for _, operator := range o.operators {
		if operator.AppliesTo(value) {
			return true
		}
	}

	return false
}
