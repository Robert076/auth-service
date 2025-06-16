package user_service

import (
	"fmt"
	"net/http"

	"github.com/Robert076/auth-service/internal/user"
)

func IsValidHttpRequest(incomingRequest *http.Request, expectedMethod string) error {
	if incomingRequest.Method != expectedMethod {
		return fmt.Errorf("error checking validity of request: got %s, expected %s", incomingRequest.Method, expectedMethod)
	}
	return nil
}

func IsValidUser(user user.RegisterUserDTO) error {
	if user.Username == "" {
		return fmt.Errorf("username cannot be empty")
	}

	if user.Email == "" {
		return fmt.Errorf("email cannot be empty")
	}

	if user.Password == "" {
		return fmt.Errorf("password cannot be empty")
	}

	return nil
}
