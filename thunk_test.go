package thunk

import (
	"errors"
	"reflect"
	"testing"
)

func TestRunSafelyWith(t *testing.T) {
	tests := []struct {
		name                   string
		thunk                  func()
		expectedCallbackResult error
	}{
		{
			"default",
			func() {
				i := 0
				i++
			},
			nil,
		},
		{
			"error",
			func() {
				panic(errors.New("zalgo"))
			},
			errors.New("zalgo"),
		},
		{
			"panic with string",
			func() {
				panic("zalgo")
			},
			errors.New("zalgo"),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			var result error
			RunSafelyWith(test.thunk, func(err error) {
				result = err
			})
			if !reflect.DeepEqual(result, test.expectedCallbackResult) {
				t.Errorf("Expected %v, got %v", test.expectedCallbackResult, result)
			}
		})
	}
}

func TestRunSafely(t *testing.T) {
	tests := []struct {
		name  string
		thunk func()
	}{
		{
			"default",
			func() {
				i := 0
				i++
			},
		},
		{
			"panic",
			func() {
				panic("zalgo")
			},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			RunSafely(test.thunk)
		})
	}
}
