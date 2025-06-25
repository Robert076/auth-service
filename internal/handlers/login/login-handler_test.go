package login_handler

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Robert076/auth-service/internal/user"
	"github.com/Robert076/auth-service/mocks"
	"github.com/stretchr/testify/assert"
)

func TestLogin(t *testing.T) {
	type testCase struct {
		name           string
		method         string
		body           interface{}
		mockSetup      func(repo *mocks.IRepository)
		expectedStatus int
	}

	validUser := user.LoginUserDTO{
		Username: "validusername",
		Email:    "valid@email.com",
		Password: "validpassword123@$",
	}

	tests := []testCase{
		{
			name:   "Valid request",
			method: http.MethodGet,
			body:   validUser,
			mockSetup: func(repo *mocks.IRepository) {
				repo.On("LoginUser", validUser).Return(nil)
			},
			expectedStatus: http.StatusOK,
		},
		{
			name:           "Invalid request (expected GET, got POST)",
			method:         http.MethodPost,
			body:           validUser,
			mockSetup:      func(repo *mocks.IRepository) {},
			expectedStatus: http.StatusBadRequest,
		},
		{
			name:           "Invalid request (expected GET, got PUT)",
			method:         http.MethodPut,
			body:           validUser,
			mockSetup:      func(repo *mocks.IRepository) {},
			expectedStatus: http.StatusBadRequest,
		},
		{
			name:           "Invalid request (expected GET, got PATCH)",
			method:         http.MethodPatch,
			body:           validUser,
			mockSetup:      func(repo *mocks.IRepository) {},
			expectedStatus: http.StatusBadRequest,
		},
		{
			name:           "Invalid request (expected GET, got DELETE)",
			method:         http.MethodDelete,
			body:           validUser,
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
			case user.LoginUserDTO:
				bodyBytes, _ = json.Marshal(v)
				reader = bytes.NewReader(bodyBytes)
			case string:
				reader = bytes.NewReader([]byte(v))
			case nil:
				reader = bytes.NewReader([]byte{})
			default:
				t.Fatalf("unsupported body type: %T", tc.body)
			}

			req := httptest.NewRequest(tc.method, "/login", reader)
			rec := httptest.NewRecorder()

			handler := LoginHandler(mockRepo)
			handler.ServeHTTP(rec, req)

			assert.Equal(t, tc.expectedStatus, rec.Code)
		})
	}
}
