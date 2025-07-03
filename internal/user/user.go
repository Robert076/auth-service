package user

import "time"

type RegisterUserDTO struct {
	Id        int       `json:"id,omitempty"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"createdAt,omitempty"`
}

type LoginUserDTO struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type AuthorizeUserDTO struct {
	Email string `json:"email"`
}

type UserDTO struct {
	Id           int       `json:"id,omitempty"`
	Username     string    `json:"username"`
	Email        string    `json:"email"`
	Password     string    `json:"password"`
	SessionToken string    `json:"sessionToken"`
	CsrfToken    string    `json:"csrfToken"`
	CreatedAt    time.Time `json:"createdAt,omitempty"`
}
