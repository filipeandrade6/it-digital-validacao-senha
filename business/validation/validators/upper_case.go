package validators

var ValidUpperCase = map[rune]void{
	97: v, 98: v, 99: v, 100: v, 101: v, 102: v, 103: v, 104: v, 105: v, 106: v,
	107: v, 108: v, 109: v, 110: v, 111: v, 112: v, 113: v, 114: v, 115: v, 116: v,
	117: v, 118: v, 119: v, 120: v, 121: v, 122: v,
}

func hasUpperCase(s string) bool {
	for _, char := range s {
		if _, ok := ValidUpperCase[char]; ok {
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

	h.Next.Check(s)

	return true
}

func (h *HasUpperCaseValidater) SetNext(next validater) {
	h.Next = next
}
