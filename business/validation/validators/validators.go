package validators

type validater interface {
	Check(s string) bool
	SetNext(validater)
}

// sugar for struct{}
type void struct{}

var v void

type Options func()

// Go maps containing the ASCII decimal representation of the valid characters.

// numbers: 0-9 => 48-57.
var validNumber = map[rune]void{
	48: v, 49: v, 50: v, 51: v, 52: v, 53: v, 54: v, 55: v, 56: v, 57: v,
}

// symbols: ![33], @[64], #[35], $[36], %[37], ^[94], &[38], *[42], ([40], )[41], -[45], +[43].
var validSymbol = map[rune]void{
	33: v, 64: v, 35: v, 36: v, 37: v, 94: v, 38: v, 42: v, 40: v, 41: v,
	45: v, 43: v,
}

// lower-case letters: a-z => 97-122.
var validLowerCase = map[rune]void{
	65: v, 66: v, 67: v, 68: v, 69: v, 70: v, 71: v, 72: v, 73: v, 74: v,
	75: v, 76: v, 77: v, 78: v, 79: v, 80: v, 81: v, 82: v, 83: v, 84: v,
	85: v, 86: v, 87: v, 88: v, 89: v, 90: v,
}

// upper-case letters: A-Z => 65-90.
var validUpperCase = map[rune]void{
	97: v, 98: v, 99: v, 100: v, 101: v, 102: v, 103: v, 104: v, 105: v, 106: v,
	107: v, 108: v, 109: v, 110: v, 111: v, 112: v, 113: v, 114: v, 115: v, 116: v,
	117: v, 118: v, 119: v, 120: v, 121: v, 122: v,
}

// valid characters (merged from all above).
var validCharacters map[rune]void = map[rune]void{
	33: v, 35: v, 36: v, 37: v, 38: v, 40: v, 41: v, 42: v, 43: v, 45: v,
	48: v, 49: v, 50: v, 51: v, 52: v, 53: v, 54: v, 55: v, 56: v, 57: v,
	64: v, 65: v, 66: v, 67: v, 68: v, 69: v, 70: v, 71: v, 72: v, 73: v,
	74: v, 75: v, 76: v, 77: v, 78: v, 79: v, 80: v, 81: v, 82: v, 83: v,
	84: v, 85: v, 86: v, 87: v, 88: v, 89: v, 90: v, 94: v, 97: v, 98: v,
	99: v, 100: v, 101: v, 102: v, 103: v, 104: v, 105: v, 106: v, 107: v, 108: v,
	109: v, 110: v, 111: v, 112: v, 113: v, 114: v, 115: v, 116: v, 117: v, 118: v,
	119: v, 120: v, 121: v, 122: v,
}
