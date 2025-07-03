package authorize_handler

import (
	"log"
	"net/http"

	"github.com/Robert076/auth-service/internal/db/repository"
	validation_service "github.com/Robert076/auth-service/internal/service/validation-service"
)

func AuthorizeHandler(repo repository.IRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := validation_service.IsValidHttpRequest(r, http.MethodPost); err != nil {
			http.Error(w, "This endpoint only accepts POST requests", http.StatusBadRequest)
			log.Printf("This endpoint only accepts POST requests: %v", err)
			return
		}

	}
}
