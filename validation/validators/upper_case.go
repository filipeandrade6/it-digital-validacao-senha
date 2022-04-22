package validators

func hasUpperCase(s string) bool {
	for _, char := range s {
		if _, ok := validUpperCase[char]; ok {
			return true
		}
	}

	return false
}

func HasUpperCaseMW(f func(s string) bool) func(s string) bool {
	return func(s string) bool {
		if !hasUpperCase(s) {
			return false
		}

		return f(s)
	}
}
