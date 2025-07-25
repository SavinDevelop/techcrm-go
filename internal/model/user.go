package model

type User struct {
	ID        string `json:"id"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	IsActive  bool   `json:"is_active"`
	CreatedAt string `json:"created_at"`
	UpdateAt  string `json:"update_at"`
}

type CreateUser struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"min=6,max=100"`
	IsActive bool   `json:"is_active"`
}
