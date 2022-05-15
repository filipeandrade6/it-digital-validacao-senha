package tests

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/filipeandrade6/iti-digital-desafio-backend/handlers"
	v1 "github.com/filipeandrade6/iti-digital-desafio-backend/handlers/v1"
	"github.com/filipeandrade6/iti-digital-desafio-backend/validation/validators"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func TestValidateHandler(t *testing.T) {
	router := handlers.NewAPI(zap.NewNop())

	tests := []struct {
		name        string
		payload     string
		wantCode    int
		wantMessage string
	}{
		{"wrong request payload", "{\"wrong\": \"AbTp9!fok\"}/", 400, "{\"is_valid\":false}"},
		{"no password", "{\"password\": \"\"}/", 400, "{\"is_valid\":false}"},
		{"invalid length", "{\"password\": \"abcDEF1!\"}/", 400, "{\"is_valid\":false}"},
		{"missing number", "{\"password\": \"abcdefghiA!\"}/", 400, "{\"is_valid\":false}"},
		{"missing symbol", "{\"password\": \"abcdefghiA1\"}/", 400, "{\"is_valid\":false}"},
		{"missing lower case", "{\"password\": \"ABCDEFGHI1!\"}/", 400, "{\"is_valid\":false}"},
		{"missing upper case", "{\"password\": \"abcdefghi1!\"}/", 400, "{\"is_valid\":false}"},
		{"repeated character", "{\"password\": \"abcdefghiA1!a\"}/", 400, "{\"is_valid\":false}"},
		{"valid password", "{\"password\": \"AbTp9!fok\"}/", 200, "{\"is_valid\":true}"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			req := httptest.NewRequest(http.MethodPost, "/v1/validate", strings.NewReader(tt.payload))
			router.ServeHTTP(w, req)
			res := w.Result()

			message := w.Body.String()

			if w.Code != tt.wantCode || message != tt.wantMessage {
				t.Errorf("Expected %#v and %#v got %#v and %#v", tt.wantCode, tt.wantMessage, w.Code, message)
			}

			if res != nil {
				res.Body.Close()
			}

		})
	}
}

func TestCustomValidateHandler(t *testing.T) {
	router := gin.New()
	router.Use(gin.Recovery())

	// Validate only unique characteres and lower-cased letters.
	vld := v1.NewValidator(zap.NewNop(), validators.HasUniqueChars, validators.HasLowerCase)
	router.POST("/v1/validate", vld.ValidateHandle)

	tests := []struct {
		name        string
		payload     string
		wantCode    int
		wantMessage string
	}{
		{"wrong request payload", "{\"wrong\": \"AbTp9!fok\"}/", 400, "{\"is_valid\":false}"},
		{"no password", "{\"password\": \"\"}/", 400, "{\"is_valid\":false}"},
		{"invalid length", "{\"password\": \"abcDEF1!\"}/", 200, "{\"is_valid\":true}"},
		{"missing number", "{\"password\": \"abcdefghiA!\"}/", 200, "{\"is_valid\":true}"},
		{"missing symbol", "{\"password\": \"abcdefghiA1\"}/", 200, "{\"is_valid\":true}"},
		{"missing lower case", "{\"password\": \"ABCDEFGHI1!\"}/", 400, "{\"is_valid\":false}"},
		{"missing upper case", "{\"password\": \"abcdefghi1!\"}/", 200, "{\"is_valid\":true}"},
		{"repeated character", "{\"password\": \"abcdefghiA1!a\"}/", 400, "{\"is_valid\":false}"},
		{"valid password", "{\"password\": \"AbTp9!fok\"}/", 200, "{\"is_valid\":true}"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			req := httptest.NewRequest(http.MethodPost, "/v1/validate", strings.NewReader(tt.payload))
			router.ServeHTTP(w, req)
			res := w.Result()

			message := w.Body.String()

			if w.Code != tt.wantCode || message != tt.wantMessage {
				t.Errorf("Expected %#v and %#v got %#v and %#v", tt.wantCode, tt.wantMessage, w.Code, message)
			}

			if res != nil {
				res.Body.Close()
			}

		})
	}
}
