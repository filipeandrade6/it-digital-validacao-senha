package validation

import (
	"github.com/filipeandrade6/iti-digital-desafio-backend/validation/validators"
)

// NewValidator creates a function validator. The validation will be executed in the
// order they are provided.
func NewValidator(mws ...func(func(string) bool) func(string) bool) func(string) bool {
	n := func(string) bool { return true }

	// Loop backwards through the middleware provided ensuring that
	// the first middleware of the slice is the first to be executed.
	for i := (len(mws) - 1); i >= 0; i-- {
		n = mws[i](n)
	}

	return n
}

// NewDefaultValidator creates a function with the default validators for password
// including length, unique characters, number, symbol, lower and upper cased letters
// and valid characters in that order.
func NewDefaultValidator() func(string) bool {
	vlds := []func(func(string) bool) func(string) bool{
		validators.HasValidLength,
		validators.HasUniqueChars,
		validators.HasNumber,
		validators.HasSymbol,
		validators.HasLowerCase,
		validators.HasUpperCase,
		validators.HasValidChars,
	}

	return NewValidator(vlds...)
}
