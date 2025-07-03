package logout_handler

import (
	"net/http"

	"github.com/Robert076/auth-service/internal/db/repository"
	validation_service "github.com/Robert076/auth-service/internal/service/validation-service"
)

func LogoutHandler(repo repository.IRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := validation_service.IsValidHttpRequest(r, http.MethodPost); err != nil {
			http.Error(w, "This endpoint only accepts POST requests", http.StatusMethodNotAllowed)
		}
	}
}
