package operator

type LessThan struct {
	value interface{}
}

func NewLessThan(value interface{}) *LessThan {
	return &LessThan{
		value: value,
	}
}

func (l *LessThan) AppliesTo(value interface{}) bool {
	intVal, isInt := l.value.(int)
	intValGiven, isIntGiven := value.(int)

	if isInt && isIntGiven {
		return intValGiven < intVal
	}

	float32Val, isFloat32 := l.value.(float32)
	float32ValGiven, isFloat32Given := value.(float32)

	if isFloat32 && isFloat32Given {
		return float32ValGiven < float32Val
	}

	return false
}
