package validators

func hasValidLength(s string) bool {
	return (len(s) >= 9)
}

type HasValidLengthValidater struct {
	Next validater
}

func (h *HasValidLengthValidater) Check(s string) bool {
	if !hasValidLength(s) {
		return false
	}

	if h.Next != nil {
		return h.Next.Check(s)
	}

	return true
}

func (h *HasValidLengthValidater) SetNext(next validater) {
	h.Next = next
}
