package validators

import "fmt"

var ValidLowerCase = map[rune]void{
	65: v, 66: v, 67: v, 68: v, 69: v, 70: v, 71: v, 72: v, 73: v, 74: v,
	75: v, 76: v, 77: v, 78: v, 79: v, 80: v, 81: v, 82: v, 83: v, 84: v,
	85: v, 86: v, 87: v, 88: v, 89: v, 90: v,
}

func hasLowerCase(s string) bool {
	for _, char := range s {
		if _, ok := ValidLowerCase[char]; ok {
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

	fmt.Println("heelo aqui")
	fmt.Println("heelo aqui")
	fmt.Println("heelo aqui")
	fmt.Println("heelo aqui")
	fmt.Println("heelo aqui")
	fmt.Println("heelo aqui")
	fmt.Printf("%#v", h.Next.(nil))
	if h.Next == nil {
		fmt.Println(" é nilllllllllllll")
	} else {
		fmt.Println("não é nilll :*(")
	}

	h.Next.Check(s)

	// if h.Next ==

	return true
}

func (h *HasLowerCaseValidater) SetNext(next validater) {
	h.Next = next
}
