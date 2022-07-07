package storage

import (
	"github.com/sarjon/toggler"
	"testing"
)

func TestNewInMemory(t *testing.T) {
	tests := []struct {
		Toggles       []*toggler.Toggle
		ExpectedError error
	}{
		{
			Toggles: []*toggler.Toggle{
				toggler.NewToggle("x"),
				toggler.NewToggle("y"),
			},
			ExpectedError: nil,
		},
		{
			Toggles: []*toggler.Toggle{
				toggler.NewToggle("x"),
				toggler.NewToggle("y"),
				toggler.NewToggle("x"),
			},
			ExpectedError: ErrDuplicateToggle,
		},
	}

	for _, test := range tests {
		inMemory, err := NewInMemory(test.Toggles)

		if test.ExpectedError != nil && err != test.ExpectedError {
			t.Errorf("expected error '%s', but received %v", test.ExpectedError, err)

			continue
		}

		if test.ExpectedError == nil && inMemory == nil {
			t.Errorf("InMemory storage was not constructed")
		}
	}
}
