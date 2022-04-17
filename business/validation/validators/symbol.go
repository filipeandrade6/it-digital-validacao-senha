package validators

func hasSymbol(s string) bool {
	for _, char := range s {
		if _, ok := validSymbol[char]; ok {
			return true
		}
	}

	return false
}

type HasSymbolValidater struct {
	Next validater
}

func (h *HasSymbolValidater) Check(s string) bool {
	if !hasSymbol(s) {
		return false
	}

	if h.Next != nil {
		return h.Next.Check(s)
	}

	return true
}

func (h *HasSymbolValidater) SetNext(next validater) {
	h.Next = next
}
