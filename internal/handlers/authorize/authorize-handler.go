package authorize_handler

import (
	"net/http"

	"github.com/Robert076/auth-service/internal/db/repository"
)

func AuthorizeHandler(repo repository.IRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}
