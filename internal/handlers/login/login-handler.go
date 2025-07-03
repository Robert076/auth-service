package login_handler

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/Robert076/auth-service/internal/constants"
	"github.com/Robert076/auth-service/internal/db/repository"
	token_service "github.com/Robert076/auth-service/internal/service/token-service"
	validation_service "github.com/Robert076/auth-service/internal/service/validation-service"
	"github.com/Robert076/auth-service/internal/user"
)

func LoginHandler(repo repository.IRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := validation_service.IsValidHttpRequest(r, http.MethodGet); err != nil {
			http.Error(w, "Invalid method for request. This endpoint only accepts GET.", http.StatusMethodNotAllowed)
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

		if err := repo.LoginUser(u); err != nil {
			http.Error(w, "Could not get user from db", http.StatusBadRequest)
			log.Printf("%s: Could not get user from db: %v", constants.ServiceName, err)
			return
		}

		sessionToken := token_service.GenerateToken(32)
		csrfToken := token_service.GenerateToken(32)

		http.SetCookie(w, &http.Cookie{
			Name:     "session_token",
			Value:    sessionToken,
			Expires:  time.Now().Add(time.Hour * 24),
			HttpOnly: true,
		})

		http.SetCookie(w, &http.Cookie{
			Name:     "csrf_token",
			Value:    csrfToken,
			Expires:  time.Now().Add(time.Hour * 24),
			HttpOnly: false, // accessible on client side
		})

		if err := repo.SetSessionToken(u.Email, sessionToken); err != nil {
			http.Error(w, "Cannot set session token.", http.StatusBadRequest)
			log.Printf("%s: Could not set session token: %v", constants.ServiceName, err)
			return
		}

		if err := repo.SetCsrfToken(u.Email, csrfToken); err != nil {
			http.Error(w, "Cannot set csrf token.", http.StatusBadRequest)
			log.Printf("%s: Could not set csrf token: %v", constants.ServiceName, err)
			return
		}

		w.WriteHeader(http.StatusOK)
		log.Printf("User %s successfully logged in", u.Username)
	}
}
