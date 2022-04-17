package main

/*
var possibleValues map[rune]bool = map[rune]bool{
	33: false, 35: false, 36: false, 37: false, 38: false, 40: false, 41: false, 42: false, 43: false, 45: false,
	48: false, 49: false, 50: false, 51: false, 52: false, 53: false, 54: false, 55: false, 56: false, 57: false,
	64: false, 65: false, 66: false, 67: false, 68: false, 69: false, 70: false, 71: false, 72: false, 73: false,
	74: false, 75: false, 76: false, 77: false, 78: false, 79: false, 80: false, 81: false, 82: false, 83: false,
	84: false, 85: false, 86: false, 87: false, 88: false, 89: false, 90: false, 94: false, 97: false, 98: false,
	99: false, 100: false, 101: false, 102: false, 103: false, 104: false, 105: false, 106: false, 107: false, 108: false,
	109: false, 110: false, 111: false, 112: false, 113: false, 114: false, 115: false, 116: false, 117: false, 118: false,
	119: false, 120: false, 121: false, 122: false,
}
*/

// func newAcceptable() map[rune]bool {
// 	out := make(map[rune]bool, possibleValues)

// 	// possible symbols maps
// 	possible := map[rune]bool{33: false, 64: false, 35: false, 36: false, 37: false, 94: false, 38: false, 42: false, 40: false, 41: false, 45: false, 43: false}

// 	// possible numbers
// 	for i := rune(48); i <= 57; i++ {
// 		out[i] = false
// 	}

// 	// possible lower-case character
// 	for i := rune(65); i <= 90; i++ {
// 		out[i] = false
// 	}

// 	// possible upper-case character
// 	for i := rune(97); i <= 122; i++ {
// 		out[i] = false
// 	}

// 	// possible symbols
// 	for k, v := range possible {
// 		out[k] = v
// 	}

// 	return out
// }
