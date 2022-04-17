package validators

var validSymbol = map[rune]void{
	33: v, 64: v, 35: v, 36: v, 37: v, 94: v, 38: v, 42: v, 40: v, 41: v,
	45: v, 43: v,
}

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

	h.Next.Check(s)

	return true
}

func (h *HasSymbolValidater) SetNext(next validater) {
	h.Next = next
}
