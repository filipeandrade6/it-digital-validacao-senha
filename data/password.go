package data

type Password struct {
	Password string `json:"password" binding:"required"`
}
