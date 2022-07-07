package operator

import "hash/crc32"

type Percentage struct {
	percent int
}

func NewPercentage(percent int) *Percentage {
	return &Percentage{
		percent: percent,
	}
}

func (p *Percentage) AppliesTo(value interface{}) bool {
	v, ok := value.(string)
	if !ok {
		return false
	}

	checksum := crc32.Checksum([]byte(v), crc32.IEEETable)

	mod := checksum % 100

	return p.percent > int(mod)
}
