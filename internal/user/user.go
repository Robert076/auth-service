package user

import "time"

type RegisterUserDTO struct {
	Id        int       `json:"id,omitempty"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"createdAt,omitempty"`
}
