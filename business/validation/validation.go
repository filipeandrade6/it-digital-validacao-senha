package validation

import (
	"fmt"

	"github.com/filipeandrade6/iti-digital-desafio-backend/business/validation/validators"
)

func IsValid(s string) bool {
	uc := &validators.HasUpperCaseValidater{}
	lc := &validators.HasLowerCaseValidater{}
	lc.SetNext(uc)
	sy := &validators.HasSymbolValidater{}
	sy.SetNext(lc)
	nu := &validators.HasNumberValidater{}
	nu.SetNext(sy)
	vc := &validators.HasValidCharactersValidater{}
	vc.SetNext(nu)
	un := &validators.HasUniqueCharactersValidater{}
	un.SetNext(vc)
	le := &validators.HasValidLengthValidater{}
	le.SetNext(un)

	fmt.Println()

	return le.Check(s)
}
