package validators

func hasValidLength(s string) bool {
	return (len(s) >= 9)
}

// HasValidLength checks if the provided string has 9 or more characters.
func HasValidLength(f func(s string) bool) func(s string) bool {
	return func(s string) bool {
		if !hasValidLength(s) {
			return false
		}

		return f(s)
	}
}
