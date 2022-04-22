package validators

import "fmt"

func hasSymbol(s string) bool {
	fmt.Println("...hasSymbol")
	for _, char := range s {
		if _, ok := validSymbol[char]; ok {
			return true
		}
	}

	return false
}

func HasSymbolMW(f func(s string) bool) func(s string) bool {
	return func(s string) bool {
		if !hasSymbol(s) {
			return false
		}

		return f(s)
	}
}
