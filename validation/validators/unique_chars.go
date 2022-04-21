package validators

func hasUniqueChars(s string) bool {
	var charsInString map[rune]struct{} = make(map[rune]struct{}, len(s))

	for _, char := range s {
		if _, alreadyExist := charsInString[char]; alreadyExist {
			return false
		}

		charsInString[char] = struct{}{}
	}

	return true
}

func HasUniqueCharactersMW(f func(s string) bool) func(s string) bool {
	return func(s string) bool {
		if !hasUniqueChars(s) {
			return false
		}

		return f(s)
	}
}
