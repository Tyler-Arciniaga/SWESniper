package models

type User struct {
	Id    string `json:"id"`
	Email string `json:"email"`
}

type SignUpUser struct {
	Email        string `json:"email"`
	Password     string `json:"password"`
	EmailConfirm bool   `json:"email_confirm"`
}
