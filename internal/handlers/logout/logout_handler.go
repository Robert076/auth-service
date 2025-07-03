package logout_handler

import (
	"net/http"

	"github.com/Robert076/auth-service/internal/db/repository"
)

func LogoutHandler(repo repository.IRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}
