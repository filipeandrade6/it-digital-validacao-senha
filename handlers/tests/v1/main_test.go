package tests

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/filipeandrade6/iti-digital-desafio-backend/handlers"
	"go.uber.org/zap"
)

func TestValidateHandler(t *testing.T) {
	router := handlers.NewAPI(zap.NewNop())
	w := httptest.NewRecorder()

	// payload := strings.NewReader("{\"password\": \"1234\"}/")
	// req := httptest.NewRequest(http.MethodGet, "/v1/validate", payload)
	// router.ServeHTTP(w, req)

	tests := []struct {
		name        string
		payload     string
		wantCode    int
		wantMessage string
	}{
		{"wrong request payload", "{\"wrong\": \"AbTp9!fok\"}/", 400, "{\"status\":\"invalid payload\"}"},
		{"no password", "{\"password\": \"\"}/", 400, "{\"status\":\"invalid password\"}"},
		{"invalid length", "{\"password\": \"abcDEF1!\"}/", 400, "{\"status\":\"invalid password\"}"},
		{"missing number", "{\"password\": \"abcdefghiA!\"}/", 400, "{\"status\":\"invalid password\"}"},
		{"missing symbol", "{\"password\": \"abcdefghiA1\"}/", 400, "{\"status\":\"invalid password\"}"},
		{"missing lower case", "{\"password\": \"ABCDEFGHI1!\"}/", 400, "{\"status\":\"invalid password\"}"},
		{"missing upper case", "{\"password\": \"abcdefghi1!\"}/", 400, "{\"status\":\"invalid password\"}"},
		{"repeated character", "{\"password\": \"abcdefghiA1!a\"}/", 400, "{\"status\":\"invalid password\"}"},
		{"valid password", "{\"password\": \"AbTp9!fok\"}/", 200, "{\"status\":\"good to go :)\"}"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodPost, "/v1/validate", strings.NewReader(tt.payload))
			router.ServeHTTP(w, req)
			res := w.Result()
			if res != nil {
				defer res.Body.Close()
			}

			body, err := ioutil.ReadAll(res.Body)
			if err != nil {
				t.Errorf("expected error to be nil when reading body got %v", err)
			}

			if err != nil {
				t.Errorf("expected error to be nil got %v", err)
			}

			message := string(data)

			if w.Code != tt.wantCode || message != tt.wantMessage {
				t.Errorf("Expected %#v and %#v got %#v and %#v", tt.wantCode, tt.wantMessage, w.Code, message)
			}
		})
	}

}
