package authorize_handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/Robert076/auth-service/internal/db/repository"
	validation_service "github.com/Robert076/auth-service/internal/service/validation-service"
	"github.com/Robert076/auth-service/internal/user"
)

func AuthorizeHandler(repo repository.IRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := validation_service.IsValidHttpRequest(r, http.MethodPost); err != nil {
			http.Error(w, "This endpoint only accepts POST requests", http.StatusBadRequest)
			log.Printf("This endpoint only accepts POST requests: %v", err)
			return
		}
		var u user.AuthorizeUserDTO

		if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
			http.Error(w, "Error decoding body", http.StatusBadRequest)
			log.Printf("Error decoding JSON body for authorization: %v", err)
			return
		}

		if err := repo.GetUserByEmail(u.Email); err != nil {
			http.Error(w, "Error retrieving user from db", http.StatusBadRequest)
			log.Printf("Error retrieving user from db: %v", err)
			return
		}
	}
}
