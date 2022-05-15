package validation

import (
	"testing"

	"github.com/filipeandrade6/iti-digital-desafio-backend/validation/validators"
)

func TestDefaultValidator(t *testing.T) {
	v := NewDefaultValidator()

	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"empty password", args{""}, false},
		{"two characters with duplicate", args{"aa"}, false},
		{"two distinct characters", args{"ab"}, false},
		{"upper and lower-case repeated characters", args{"AAAbbbCc"}, false},
		{"repeated characters with valid lenght", args{"AbTp9!foo"}, false},
		{"repeated upper cased character", args{"AbTp9!foA"}, false},
		{"password with space", args{"AbTp9 fok"}, false},
		{"valid password", args{"AbTp9!fok"}, true},
		{"password without number", args{"AbTpX!fok"}, false},
		{"password without upper case", args{"abtp9!fok"}, false},
		{"password without lower case", args{"ABTP9!FOK"}, false},
		{"password without symbol", args{"AbTp9Xfok"}, false},
		{"password with 8 characters", args{"AbTp9!fo"}, false},
		{"password with 10 characters", args{"AbTp9!fok1"}, true},
		{"password with accent", args{"âbTp9!fok"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := v(tt.args.s); got != tt.want {
				t.Errorf("isValid() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEmptyValidator(t *testing.T) {
	v := NewValidator()

	if !v("") {
		t.Errorf("Should return true for empty password without any validators.")
	}

	if !v("hello") {
		t.Errorf("Should return true for empty password without any validators.")
	}
}

func TestOnlyValidLengthValidator(t *testing.T) {
	v := NewValidator(validators.HasValidLength)

	if v("123") {
		t.Errorf("Should return false for invalid length password with only HasValidLength Validator.")
	}

	if !v("123456789") {
		t.Errorf("Should return true for invalid length password with only HasValidLength Validator.")
	}
}

func TestOnlyLowerCaseValidator(t *testing.T) {
	v := NewValidator(validators.HasLowerCase)

	if v("ABD") {
		t.Errorf("Should return false for password without lower cased letter with only HasLowerCase Validator.")
	}

	if !v("aABD") {
		t.Errorf("Should return true for password with lower cased letter with only HasLowerCase Validator.")
	}
}

func TestOnlyNumberValidator(t *testing.T) {
	v := NewValidator(validators.HasNumber)

	if v("abc") {
		t.Errorf("Should return false for password without number with only HasNumber Validator.")
	}

	if !v("123") {
		t.Errorf("Should return true for password with number with only HasNumber Validator.")
	}
}

func TestOnlySymbolValidator(t *testing.T) {
	v := NewValidator(validators.HasSymbol)

	if v("abc123A") {
		t.Errorf("Should return false for password without symbol with only HasSymbol Validator.")
	}

	if !v("!") {
		t.Errorf("Should return true for password with symbol with only Number Validator.")
	}
}

func TestOnlyUniqueCharsValidator(t *testing.T) {
	v := NewValidator(validators.HasUniqueChars)

	if v("121234") {
		t.Errorf("Should return false for password with repeated character with only HasUniqueChars Validator.")
	}

	if !v("123") {
		t.Errorf("Should return true for password without repeated character with only HasUniqueChars Validator.")
	}
}

func TestOnlyUpperCaseValidator(t *testing.T) {
	v := NewValidator(validators.HasUpperCase)

	if v("abc123") {
		t.Errorf("Should return false for password without upper cased letter with only HasUpperCase Validator.")
	}

	if !v("A123") {
		t.Errorf("Should return true for password with upper cased letter with only HasUpperCase Validator.")
	}
}

func TestOnlyValidCharsValidator(t *testing.T) {
	v := NewValidator(validators.HasValidChars)

	if v("?123") {
		t.Errorf("Should return false for password with invalid characters with only HasValidChars Validator.")
	}

	if !v("abcABC!@#$%&*()-+123") {
		t.Errorf("Should return true for password with valid characters with only HasValidChars Validator.")
	}
}
