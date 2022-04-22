package data

// Password is a binding struct for the provided JSON request.
type Password struct {
	Password string `json:"password" binding:"required"`
}
