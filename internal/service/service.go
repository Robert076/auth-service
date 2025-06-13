package service

import (
	"fmt"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

func IsValidHttpRequest(incomingRequest *http.Request, expectedMethod string) error {
	if incomingRequest.Method != expectedMethod {
		return fmt.Errorf("error checking validity of request: got %s, expected %s", incomingRequest.Method, expectedMethod)
	}
	return nil
}

func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("error hashing password: %v", err)
	}
	return string(hashedPassword), nil
}

func CompareHash(hash string, password string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password)) == nil
}
