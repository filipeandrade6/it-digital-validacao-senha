package validators

import "fmt"

func hasValidChars(s string) bool {
	fmt.Println("...hasValidChars")
	for _, char := range s {
		if _, ok := validCharacters[char]; !ok {
			return false
		}
	}

	return true
}

func HasValidCharsMW(f func(s string) bool) func(s string) bool {
	return func(s string) bool {
		if !hasValidChars(s) {
			return false
		}

		return f(s)
	}
}
