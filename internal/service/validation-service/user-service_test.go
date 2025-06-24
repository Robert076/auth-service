package user_service

import (
	"net/http"
	"testing"
	"time"

	"github.com/Robert076/auth-service/internal/user"
)

func TestIsValidHttpRequest(t *testing.T) {
	tests := []struct {
		name           string
		actualMethod   string
		expectedMethod string
		shouldFail     bool
	}{
		{"Valid GET", http.MethodGet, http.MethodGet, false},
		{"Valid POST", http.MethodPost, http.MethodPost, false},
		{"Valid PUT", http.MethodPut, http.MethodPut, false},
		{"Valid DELETE", http.MethodDelete, http.MethodDelete, false},
		{"Invalid mismatch", http.MethodPost, http.MethodGet, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req, err := http.NewRequest(tt.actualMethod, "http://localhost:8080", nil)
			if err != nil {
				t.Fatalf("Failed to create request: %v", err)
			}

			err = IsValidHttpRequest(req, tt.expectedMethod)
			if tt.shouldFail && err == nil {
				t.Errorf("Expected failure but got none")
			}
			if !tt.shouldFail && err != nil {
				t.Errorf("Expected success but got error: %v", err)
			}
		})
	}
}

func TestIsValidUserRegister(t *testing.T) {
	baseUser := user.RegisterUserDTO{
		Id:        1,
		Username:  "mockusername",
		Email:     "mockemail@email.com",
		Password:  "password123$",
		CreatedAt: time.Now(),
	}

	tests := []struct {
		name    string
		modify  func(u *user.RegisterUserDTO)
		wantErr bool
	}{
		{
			name: "Missing username",
			modify: func(u *user.RegisterUserDTO) {
				u.Username = ""
			},
			wantErr: true,
		},
		{
			name: "Missing email",
			modify: func(u *user.RegisterUserDTO) {
				u.Email = ""
			},
			wantErr: true,
		},
		{
			name: "Missing password",
			modify: func(u *user.RegisterUserDTO) {
				u.Password = ""
			},
			wantErr: true,
		},
		{
			name: "Valid user",
			modify: func(u *user.RegisterUserDTO) {
				// no change
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := baseUser
			tt.modify(&u)
			err := IsValidUserRegister(u)
			if tt.wantErr && err == nil {
				t.Errorf("Expected error but got nil")
			}
			if !tt.wantErr && err != nil {
				t.Errorf("Expected no error, but got: %v", err)
			}
		})
	}
}

func TestIsValidUserLogin(t *testing.T) {
	baseUser := user.LoginUserDTO{
		Username: "username",
		Email:    "email@email.com",
		Password: "password123$",
	}

	tests := []struct {
		name    string
		modify  func(u *user.LoginUserDTO)
		wantErr bool
	}{
		{
			name: "Missing both email and username",
			modify: func(u *user.LoginUserDTO) {
				u.Username = ""
				u.Email = ""
			},
			wantErr: true,
		},
		{
			name: "Missing just email, not username too",
			modify: func(u *user.LoginUserDTO) {
				u.Email = ""
			},
			wantErr: false,
		},
		{
			name: "Missing just username, not email too",
			modify: func(u *user.LoginUserDTO) {
				u.Username = ""
			},
			wantErr: false,
		},
		{
			name: "Missing password",
			modify: func(u *user.LoginUserDTO) {
				u.Password = ""
			},
			wantErr: true,
		},
		{
			name: "Missing nothing, should work",
			modify: func(u *user.LoginUserDTO) {
				// nothing here
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := baseUser
			tt.modify(&u)
			err := IsValidUserLogin(u)
			if tt.wantErr && err == nil {
				t.Errorf("Expected error but got nil")
			}
			if !tt.wantErr && err != nil {
				t.Errorf("Expected no error, but got: %v", err)
			}
		})
	}
}
