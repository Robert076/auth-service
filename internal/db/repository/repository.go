package repository

import user "github.com/Robert076/auth-service/internal/user"

type IRepository interface {
	RegisterUser(user user.RegisterUserDTO) error
}
