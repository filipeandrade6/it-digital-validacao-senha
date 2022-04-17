package validators

func hasUpperCase(s string) bool {
	for _, char := range s {
		if _, ok := validUpperCase[char]; ok {
			return true
		}
	}

	return false
}

type HasUpperCaseValidater struct {
	Next validater
}

func (h *HasUpperCaseValidater) Check(s string) bool {
	if !hasUpperCase(s) {
		return false
	}

	if h.Next != nil {
		return h.Next.Check(s)
	}

	return true
}

func (h *HasUpperCaseValidater) SetNext(next validater) {
	h.Next = next
}
