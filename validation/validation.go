package validation

import (
	"fmt"

	"github.com/filipeandrade6/iti-digital-desafio-backend/validation/validators"
)

// ! toda vez é instanciado as estruturas - fazer para criar uma vez...

func IsValid(s string) bool {
	uc := &validators.HasUpperCaseValidater{}
	lc := &validators.HasLowerCaseValidater{}
	sy := &validators.HasSymbolValidater{}
	nu := &validators.HasNumberValidater{}
	vc := &validators.HasValidCharactersValidater{}
	un := &validators.HasUniqueCharactersValidater{}
	le := &validators.HasValidLengthValidater{}

	lc.SetNext(uc)
	sy.SetNext(lc)
	nu.SetNext(sy)
	vc.SetNext(nu)
	un.SetNext(vc)
	le.SetNext(un)

	fmt.Println()

	return le.Check(s)
}

type Validator struct{}

func NewValidator() *Validator {
	return &Validator{}
}

func (v *Validator) IsValid(s string) bool {
	return true
}
