package validators

// sugar for struct{}
type void struct{}

var (
	// sugar for struct{}{}
	v void

	// Numbers: 0-9.
	validNumber = map[rune]void{
		rune('0'): v, rune('1'): v, rune('2'): v, rune('3'): v, rune('4'): v, rune('5'): v,
		rune('6'): v, rune('7'): v, rune('8'): v, rune('9'): v,
	}

	// Symbols: !, @, #, $, %, ^, &, *, (, ), -, +.
	validSymbol = map[rune]void{
		rune('!'): v, rune('@'): v, rune('#'): v, rune('$'): v, rune('%'): v, rune('^'): v,
		rune('&'): v, rune('*'): v, rune('('): v, rune(')'): v, rune('-'): v, rune('+'): v,
	}

	// Lower-case letters: a-z.
	validLowerCase = map[rune]void{
		rune('a'): v, rune('b'): v, rune('c'): v, rune('d'): v, rune('e'): v, rune('f'): v,
		rune('g'): v, rune('h'): v, rune('i'): v, rune('j'): v, rune('k'): v, rune('l'): v,
		rune('m'): v, rune('n'): v, rune('o'): v, rune('p'): v, rune('q'): v, rune('r'): v,
		rune('s'): v, rune('t'): v, rune('u'): v, rune('v'): v, rune('w'): v, rune('x'): v,
		rune('y'): v, rune('z'): v,
	}

	// Upper-case letters: A-Z.
	validUpperCase = map[rune]void{
		rune('A'): v, rune('B'): v, rune('C'): v, rune('D'): v, rune('E'): v, rune('F'): v,
		rune('G'): v, rune('H'): v, rune('I'): v, rune('J'): v, rune('K'): v, rune('L'): v,
		rune('M'): v, rune('N'): v, rune('O'): v, rune('P'): v, rune('Q'): v, rune('R'): v,
		rune('S'): v, rune('T'): v, rune('U'): v, rune('V'): v, rune('W'): v, rune('X'): v,
		rune('Y'): v, rune('Z'): v,
	}

	// Valid characters: merge from above.
	validCharacters map[rune]void = make(map[rune]void)
)

func init() {
	for _, m := range []map[rune]void{validNumber, validSymbol, validLowerCase, validUpperCase} {
		for k, v := range m {
			validCharacters[k] = v
		}
	}
}
