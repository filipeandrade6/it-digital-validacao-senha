package validators

func hasUniqueCharacters(s string) bool {
	var charsInString map[rune]struct{} = make(map[rune]struct{}, len(s))

	for _, char := range s {
		if _, alreadyExist := charsInString[char]; alreadyExist {
			return false
		}

		charsInString[char] = struct{}{}
	}

	return true
}

type HasUniqueCharactersValidater struct {
	Next validater
}

func (h *HasUniqueCharactersValidater) Check(s string) bool {
	if !hasUniqueCharacters(s) {
		return false
	}

	if h.Next != nil {
		return h.Next.Check(s)
	}

	return true
}

func (h *HasUniqueCharactersValidater) SetNext(next validater) {
	h.Next = next
}
