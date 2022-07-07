package operator

type Operator interface {
	AppliesTo(value interface{}) bool
}
