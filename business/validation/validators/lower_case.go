package validators

func hasLowerCase(s string) bool {
	for _, char := range s {
		if _, ok := validLowerCase[char]; ok {
			return true
		}
	}

	return false
}

type HasLowerCaseValidater struct {
	Next validater
}

func (h *HasLowerCaseValidater) Check(s string) bool {
	if !hasLowerCase(s) {
		return false
	}

	if h.Next != nil {
		return h.Next.Check(s)
	}

	return true
}

func (h *HasLowerCaseValidater) SetNext(next validater) {
	h.Next = next
}
