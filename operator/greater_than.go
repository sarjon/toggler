package operator

type GreaterThan struct {
	value interface{}
}

func NewGreaterThan(value interface{}) *GreaterThan {
	return &GreaterThan{
		value: value,
	}
}

func (g *GreaterThan) AppliesTo(value interface{}) bool {
	intVal, isInt := g.value.(int)
	intValGiven, isIntGiven := value.(int)

	if isInt && isIntGiven {
		return intValGiven > intVal
	}

	float32Val, isFloat32 := g.value.(float32)
	float32ValGiven, isFloat32Given := value.(float32)

	if isFloat32 && isFloat32Given {
		return float32ValGiven > float32Val
	}

	return false
}
