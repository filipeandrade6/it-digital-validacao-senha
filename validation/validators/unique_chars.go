package validators

func hasUniqueChars(s string) bool {
	var charsInString = make(map[rune]struct{}, len(s))

	for _, char := range s {
		if _, alreadyExist := charsInString[char]; alreadyExist {
			return false
		}

		charsInString[char] = struct{}{}
	}

	return true
}

// HasUniqueChars checks if the provided string has unique characters.
func HasUniqueChars(f func(s string) bool) func(s string) bool {
	return func(s string) bool {
		if !hasUniqueChars(s) {
			return false
		}

		return f(s)
	}
}
