package validators

var ValidNumber = map[rune]void{
	48: v, 49: v, 50: v, 51: v, 52: v, 53: v, 54: v, 55: v, 56: v, 57: v,
}

func hasNumber(s string) bool {
	for _, char := range s {
		if _, ok := ValidNumber[char]; ok {
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

	h.Next.Check(s)

	return true
}

func (h *HasNumberValidater) SetNext(next validater) {
	h.Next = next
}
