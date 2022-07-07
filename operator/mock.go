package operator

type mock struct {
	value bool
}

func newMock(value bool) *mock {
	return &mock{
		value: value,
	}
}

func (m *mock) AppliesTo(value interface{}) bool {
	return m.value
}
