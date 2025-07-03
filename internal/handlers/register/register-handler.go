package register_handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/Robert076/auth-service/internal/constants"
	"github.com/Robert076/auth-service/internal/db/repository"
	validation_service "github.com/Robert076/auth-service/internal/service/validation-service"
	"github.com/Robert076/auth-service/internal/user"
)

func RegisterHandler(repo repository.IRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := validation_service.IsValidHttpRequest(r, http.MethodPost); err != nil {
			http.Error(w, "Invalid method for r. This endpoint only accepts POST.", http.StatusBadRequest)
			log.Printf("%s: Error validating r for POST (Register). The issue might be that this endpoint only accepts POST rs. Error: %v", constants.ServiceName, err)
			return
		}

		var u user.RegisterUserDTO

		if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
			http.Error(w, "Invalid JSON payload", http.StatusBadRequest)
			log.Printf("%s: Invalid JSON payload: %v", constants.ServiceName, err)
			return
		}

		if err := validation_service.IsValidUserRegister(u); err != nil {
			http.Error(w, "Invalid user received for register endpoint.", http.StatusBadRequest)
			log.Printf("%s: Invalid user received for register endpoint: %v", constants.ServiceName, err)
			return
		}

		if err := repo.RegisterUser(u); err != nil {
			http.Error(w, "Could not insert user in db", http.StatusBadRequest)
			log.Printf("%s: Could not insert user in db: %v", constants.ServiceName, err)
			return
		}

		w.WriteHeader(http.StatusOK)
		log.Printf("%s: Successfully added user in db.", constants.ServiceName)
	}
}
