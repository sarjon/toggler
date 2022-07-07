package toggler

const (
	ActiveConditionally string = "ACTIVE_CONDITIONALLY"
	ActiveAlways        string = "ACTIVE_ALWAYS"
	ActiveNever         string = "ACTIVE_NEVER"
)

type ToggleActivation string

type Toggle struct {
	name       string
	conditions []Condition
	activation string
	strategy   string
}

func NewToggle(name string, conditions ...Condition) *Toggle {
	return newToggle(name, conditions, ActiveConditionally, StrategyAffirmative)
}

func newToggle(name string, conditions []Condition, activation string, strategy string) *Toggle {
	return &Toggle{
		name:       name,
		conditions: conditions,
		activation: activation,
		strategy:   strategy,
	}
}

func (t *Toggle) Name() string {
	return t.name
}

func (t *Toggle) Strategy() string {
	return t.strategy
}

func (t *Toggle) Activation() string {
	return t.activation
}

func (t *Toggle) Conditions() []Condition {
	return t.conditions
}

func (t *Toggle) IsZero() bool {
	return t.name == ""
}

func (t *Toggle) SetActivation(activation string) {
	t.activation = activation
}

func (t *Toggle) SetStrategy(strategy string) {
	t.strategy = strategy
}
