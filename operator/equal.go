package operator

type equal struct {
	value interface{}
}

func NewEqual(value interface{}) *equal {
	return &equal{
		value: value,
	}
}

func (e *equal) AppliesTo(value interface{}) bool {
	return e.value == value
}
