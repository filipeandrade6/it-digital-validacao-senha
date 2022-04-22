package validation

import "github.com/filipeandrade6/iti-digital-desafio-backend/validation/validators"

// TODO wrap the function in inverse order to be called as the param list
func NewValidator(mws ...func(func(string) bool) func(string) bool) func(string) bool {
	n := func(string) bool { return true }

	// Loop backwards through the middleware provided ensuring that
	// the first middleware of the slice is the first to be executed.
	for i := (len(mws) - 1); i >= 0; i-- {
		n = mws[i](n)
	}

	return n
}

func NewDefaultValidator() func(string) bool {
	vlds := []func(func(string) bool) func(string) bool{
		validators.HasValidCharsMW,
		validators.HasLowerCaseMW,
		validators.HasUpperCaseMW,
		validators.HasSymbolMW,
		validators.HasNumberMW,
		validators.HasUniqueCharsMW,
		validators.HasValidLengthMW,
	}

	return NewValidator(vlds...)
}
