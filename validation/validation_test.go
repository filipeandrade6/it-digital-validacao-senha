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
}

func TestOnlyValidLengthValidator(t *testing.T) {
	v := NewValidator(validators.HasValidLength)

	if v("123") {
		t.Errorf("Should return false for invalid length password with only Valid Length Validator.")
	}
}
