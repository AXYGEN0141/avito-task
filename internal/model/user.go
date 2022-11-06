package model

type Response struct {
	Text    string `json:"text"`
	Message string `json:"message"`
	Data    interface{}
}

type User struct {
	ID       int     `json:"id"`
	Name     string  `json:"name"`
	Email    string  `json:"email"`
	Password string  `json:"password"`
	Account  Account `json:"account"`
}
