package login_handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/Robert076/auth-service/internal/constants"
	"github.com/Robert076/auth-service/internal/db/repository"
	validation_service "github.com/Robert076/auth-service/internal/service/validation-service"
	"github.com/Robert076/auth-service/internal/user"
)

func LoginHandler(repo repository.IRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := validation_service.IsValidHttpRequest(r, http.MethodGet); err != nil {
			http.Error(w, "Invalid method for request. This endpoint only accepts GET.", http.StatusBadRequest)
			log.Printf("%s: Error validating request for GET (login). The issue might be that this endpoint only accepts GET rs. Error: %v", constants.ServiceName, err)
			return
		}

		var u user.LoginUserDTO

		if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
			http.Error(w, "Invalid JSON payload", http.StatusBadRequest)
			log.Printf("%s: Invalid JSON payload: %v", constants.ServiceName, err)
			return
		}

		if err := validation_service.IsValidUserLogin(u); err != nil {
			http.Error(w, "Invalid user received for login endpoint.", http.StatusBadRequest)
			log.Printf("%s: Invalid user received for login endpoint: %v", constants.ServiceName, err)
			return
		}

	}
}
