package repository

import user "github.com/Robert076/auth-service/internal/user"

type IRepository interface {
	RegisterUser(user user.RegisterUserDTO) error

	LoginUser(user user.LoginUserDTO) error

	SetSessionToken(username string, sessionToken string) error

	SetCsrfToken(email string, csrfToken string) error

	GetUserByEmail(email string) (user.UserDTO, error)
}
