package validation

import (
	"fmt"

	"github.com/filipeandrade6/iti-digital-desafio-backend/business/validation/validators"
)

// * caracteres válidos / mapeamento de caractere para número decimal em ASCII:
// * números: 0-9 => 48-57.
// * letras maiúsculas: A-Z => 65-90.
// * letras minúsculas: a-z => 97-122.
// * simbolos: ![33], @[64], #[35], $[36], %[37], ^[94], &[38], *[42], ([40], )[41], -[45], +[43].

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

// 	if ok := validators.HasValidLength(s); !ok {
// 		return false
// 	}

// 	if ok := validators.HasUniqueCharacters(s); !ok {
// 		return false
// 	}

// 	if ok := validators.HasValidCharacters(s); !ok {
// 		return false
// 	}

// 	if ok := validators.HasNumber(s); !ok {
// 		return false
// 	}

// 	if ok := validators.HasSymbol(s); !ok {
// 		return false
// 	}

// 	if ok := validators.HasLowerCase(s); !ok {
// 		return false
// 	}

// 	if ok := validators.HasUpperCase(s); !ok {
// 		return false
// 	}

// 	return true
// }
