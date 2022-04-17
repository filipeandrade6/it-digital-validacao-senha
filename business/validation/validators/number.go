package validators

func hasNumber(s string) bool {
	for _, char := range s {
		if _, ok := validNumber[char]; ok {
			return true
		}
	}

	return false
}

type HasNumberValidater struct {
	Next validater
}

func (h *HasNumberValidater) Check(s string) bool {
	if !hasNumber(s) {
		return false
	}

	if h.Next != nil {
		return h.Next.Check(s)
	}

	return true
}

func (h *HasNumberValidater) SetNext(next validater) {
	h.Next = next
}
