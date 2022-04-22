package validators

import "fmt"

func hasValidLength(s string) bool {
	fmt.Println("...hasValidLength")
	return (len(s) >= 9)
}

func HasValidLengthMW(f func(s string) bool) func(s string) bool {
	return func(s string) bool {
		if !hasValidLength(s) {
			return false
		}

		return f(s)
	}
}
