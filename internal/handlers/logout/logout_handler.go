package logout_handler

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/Robert076/auth-service/internal/constants"
	"github.com/Robert076/auth-service/internal/db/repository"
	validation_service "github.com/Robert076/auth-service/internal/service/validation-service"
	"github.com/Robert076/auth-service/internal/user"
)

func LogoutHandler(repo repository.IRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := validation_service.IsValidHttpRequest(r, http.MethodPost); err != nil {
			http.Error(w, "This endpoint only accepts POST requests", http.StatusMethodNotAllowed)
			log.Printf("%s: This endpoint only accepts POST requests: %v", constants.ServiceName, err)
			return
		}

		http.SetCookie(w, &http.Cookie{
			Name:     "session_token",
			Value:    "",
			Expires:  time.Now().Add(-time.Hour),
			HttpOnly: true,
		})

		http.SetCookie(w, &http.Cookie{
			Name:     "csrf_token",
			Value:    "",
			Expires:  time.Now().Add(-time.Hour),
			HttpOnly: false,
		})

		var u user.AuthorizeUserDTO
		if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
			http.Error(w, "Error decoding body", http.StatusBadRequest)
			log.Printf("Error decoding JSON body for authorization: %v", err)
			return
		}

		if err := repo.ClearTokensByUserEmail(u.Email); err != nil {
			http.Error(w, "Error clearing tokens for user", http.StatusBadRequest)
			log.Printf("%s: Error clearing tokens for user %s: %v", constants.ServiceName, u.Email, err)
			return
		}

		w.WriteHeader(http.StatusOK)
		log.Printf("%s: Successfully logged out user %s", constants.ServiceName, u.Email)
	}
}
