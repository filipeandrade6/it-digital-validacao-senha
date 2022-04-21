package validators

func hasLowerCase(s string) bool {
	for _, char := range s {
		if _, ok := validLowerCase[char]; ok {
			return true
		}
	}

	return false
}

func HasValidLowerCaseMW(f func(s string) bool) func(s string) bool {
	return func(s string) bool {
		if !hasLowerCase(s) {
			return false
		}

		return f(s)
	}
}
