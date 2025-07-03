package authorize_handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/Robert076/auth-service/internal/db/repository"
	validation_service "github.com/Robert076/auth-service/internal/service/validation-service"
	"github.com/Robert076/auth-service/internal/user"
)

func AuthorizeHandler(repo repository.IRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := validation_service.IsValidHttpRequest(r, http.MethodPost); err != nil {
			http.Error(w, "This endpoint only accepts POST requests", http.StatusMethodNotAllowed)
			log.Printf("This endpoint only accepts POST requests: %v", err)
			return
		}
		var userAuthorize user.AuthorizeUserDTO

		if err := json.NewDecoder(r.Body).Decode(&userAuthorize); err != nil {
			http.Error(w, "Error decoding body", http.StatusBadRequest)
			log.Printf("Error decoding JSON body for authorization: %v", err)
			return
		}

		u, err := repo.GetUserByEmail(userAuthorize.Email)
		if err != nil {
			http.Error(w, "Error retrieving user from db", http.StatusBadRequest)
			log.Printf("Error retrieving user from db: %v", err)
			return
		}

		sessionToken, err := r.Cookie("session_token")

		if err != nil {
			http.Error(w, "Error retrieving session token from cookie", http.StatusBadRequest)
			log.Printf("Error retrieving session token from cookie: %v", err)
			return
		}

		if err := validateSessionToken(sessionToken.Value, u); err != nil {
			http.Error(w, "Error validating session token", http.StatusBadRequest)
			log.Printf("Error validating session token: %v", err)
			return
		}

		csrfToken := r.Header.Get("X-CSRF-Token")

		if err := validateCsrfToken(csrfToken, u); err != nil {
			http.Error(w, "Error validating csrf token", http.StatusBadRequest)
			log.Printf("Error validating csrf token: %v", err)
			return
		}

		w.WriteHeader(http.StatusOK)
		log.Printf("Successfully authorized user %s", u.Email)
	}
}

func validateSessionToken(sessionToken string, u user.UserDTO) error {
	if sessionToken == "" {
		return fmt.Errorf("session token cannot be empty")
	}

	if sessionToken != u.SessionToken {
		return fmt.Errorf("session token received differs from the one in the database")
	}

	return nil
}

func validateCsrfToken(csrfToken string, u user.UserDTO) error {
	if csrfToken == "" {
		return fmt.Errorf("csrf token cannot be empty")
	}

	if csrfToken != u.CsrfToken {
		return fmt.Errorf("csrf token received differs from the one in the database")
	}

	return nil
}
