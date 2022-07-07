package operator

type InSet struct {
	values []interface{}
}

func NewInSet(values []interface{}) *InSet {
	return &InSet{
		values: values,
	}
}

func (i *InSet) AppliesTo(value interface{}) bool {
	for _, v := range i.values {
		if v == value {
			return true
		}
	}

	return false
}
