// Package data provides data structure utilized in the application.
package data

// Password is a binding struct for the provided JSON request with validation.
type Password struct {
	Password string `json:"password" binding:"required"`
}
