package validators

func hasNumber(s string) bool {
	for _, char := range s {
		if _, ok := validNumber[char]; ok {
			return true
		}
	}

	return false
}

func HasNumberMW(f func(s string) bool) func(s string) bool {
	return func(s string) bool {
		if !hasNumber(s) {
			return false
		}

		return f(s)
	}
}
