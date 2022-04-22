package validators

func hasSymbol(s string) bool {
	for _, char := range s {
		if _, ok := validSymbol[char]; ok {
			return true
		}
	}

	return false
}

// HasSymbol checks if the provided string has the valid symbols.
func HasSymbol(f func(s string) bool) func(s string) bool {
	return func(s string) bool {
		if !hasSymbol(s) {
			return false
		}

		return f(s)
	}
}
