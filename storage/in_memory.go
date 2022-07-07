package storage

import (
	"errors"
	"github.com/sarjon/toggler"
)

var (
	ErrDuplicateToggle = errors.New("duplicate toggle")
)

type InMemory struct {
	toggles map[string]*toggler.Toggle
}

func NewInMemory(toggles []*toggler.Toggle) (*InMemory, error) {
	mappedByNameToggles := map[string]*toggler.Toggle{}

	for _, t := range toggles {
		_, found := mappedByNameToggles[t.Name()]
		if found {
			return nil, ErrDuplicateToggle
		}

		mappedByNameToggles[t.Name()] = t
	}

	return &InMemory{
		toggles: mappedByNameToggles,
	}, nil
}

func (i *InMemory) Get(name string) (*toggler.Toggle, bool) {
	toggle, found := i.toggles[name]

	return toggle, found
}
