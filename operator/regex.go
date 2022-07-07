package operator

import "regexp"

type Regex struct {
	rg *regexp.Regexp
}

func NewRegex(expr string) *Regex {
	rg, err := regexp.Compile(expr)
	if err != nil {
		panic(err)
	}

	return &Regex{
		rg: rg,
	}
}

func (r *Regex) AppliesTo(value interface{}) bool {
	val, ok := value.(string)
	if !ok {
		return false
	}

	matched := r.rg.Match([]byte(val))

	return matched
}
