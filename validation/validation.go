package validation

import "github.com/filipeandrade6/iti-digital-desafio-backend/validation/validators"

// length, unique, number, symbol, uc, lc, valid

type Validator struct{}

func NewValidator(mws ...func(func(string) bool) func(string) bool) func(string) bool {
	n := isValid
	for _, mw := range mws {
		n = mw(n)
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

func isValid(s string) bool {
	return true
}
