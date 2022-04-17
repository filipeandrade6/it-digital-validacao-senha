package validators

func hasValidCharacters(s string) bool {
	for _, char := range s {
		if _, ok := validCharacters[char]; !ok {
			return false
		}
	}

	return true
}

type HasValidCharactersValidater struct {
	Next validater
}

func (h *HasValidCharactersValidater) Check(s string) bool {
	if !hasValidCharacters(s) {
		return false
	}

	if h.Next != nil {
		return h.Next.Check(s)
	}

	return true
}

func (h *HasValidCharactersValidater) SetNext(next validater) {
	h.Next = next
}
