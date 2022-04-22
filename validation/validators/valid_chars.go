package validators

func hasValidChars(s string) bool {
	for _, char := range s {
		if _, ok := validCharacters[char]; !ok {
			return false
		}
	}

	return true
}

// HasValidChars checks if the provided string has valid characters.
func HasValidChars(f func(s string) bool) func(s string) bool {
	return func(s string) bool {
		if !hasValidChars(s) {
			return false
		}

		return f(s)
	}
}
