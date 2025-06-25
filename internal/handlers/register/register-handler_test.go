package register_handler

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/Robert076/auth-service/internal/user"
	"github.com/Robert076/auth-service/mocks"
	"github.com/stretchr/testify/assert"
)

func TestRegister(t *testing.T) {
	type testCase struct {
		name           string
		method         string
		body           interface{}
		mockSetup      func(repo *mocks.IRepository)
		expectedStatus int
	}

	validUser := user.RegisterUserDTO{
		Id:        1,
		Email:     "valid@email.com",
		Username:  "validusername",
		Password:  "validpassword123$#!",
		CreatedAt: time.Date(2025, time.July, 2, 2, 2, 2, 2, time.UTC),
	}

	tests := []testCase{
		{
			name:   "Valid request",
			method: http.MethodPost,
			body:   validUser,
			mockSetup: func(repo *mocks.IRepository) {
				repo.On("RegisterUser", validUser).Return(nil)
			},
			expectedStatus: http.StatusOK,
		},
		{
			name:           "Invalid request (GET not allowed)",
			method:         http.MethodGet,
			body:           validUser,
			mockSetup:      func(repo *mocks.IRepository) {},
			expectedStatus: http.StatusBadRequest,
		},
		{
			name:           "Invalid request (PUT not allowed)",
			method:         http.MethodPut,
			body:           validUser,
			mockSetup:      func(repo *mocks.IRepository) {},
			expectedStatus: http.StatusBadRequest,
		},
		{
			name:           "Invalid request (DELETE not allowed)",
			method:         http.MethodDelete,
			body:           validUser,
			mockSetup:      func(repo *mocks.IRepository) {},
			expectedStatus: http.StatusBadRequest,
		},
		{
			name:           "Invalid request (PATCH not allowed)",
			method:         http.MethodPatch,
			body:           validUser,
			mockSetup:      func(repo *mocks.IRepository) {},
			expectedStatus: http.StatusBadRequest,
		},
		{
			name:           "Invalid request (invalid body)",
			method:         http.MethodGet,
			body:           "{invalid-json",
			mockSetup:      func(repo *mocks.IRepository) {},
			expectedStatus: http.StatusBadRequest,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			mockRepo := mocks.NewIRepository(t)

			if tc.mockSetup != nil {
				tc.mockSetup(mockRepo)
			}

			var bodyBytes []byte
			var reader *bytes.Reader
			switch v := tc.body.(type) {
			case user.RegisterUserDTO:
				bodyBytes, _ = json.Marshal(v)
				reader = bytes.NewReader(bodyBytes)
			case string:
				reader = bytes.NewReader([]byte(v))
			case nil:
				reader = bytes.NewReader([]byte{})
			default:
				t.Fatalf("unsupported body type: %T", tc.body)
			}

			req := httptest.NewRequest(tc.method, "/register", reader)
			rec := httptest.NewRecorder()

			handler := RegisterHandler(mockRepo)
			handler.ServeHTTP(rec, req)

			assert.Equal(t, tc.expectedStatus, rec.Code)
		})
	}
}
