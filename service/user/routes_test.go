package user

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/akamaaru/go-forum/types"
	"github.com/gorilla/mux"
)

func TestUserServiceHandlers(t *testing.T) {
	userStore := &mockUserStore{}
	handler := NewHandler(userStore)

	t.Run("fail if user payload is invalid", func(t *testing.T) {
		payload := types.RegisterUserPayload {
			FirstName: "user",
			LastName: "smith",
			Email: "invalid",
			Password: "qwerty",
		}
		marchalled, _ := json.Marshal(payload)

		req, err := http.NewRequest(
			http.MethodPost, 
			"/register", 
			bytes.NewBuffer(marchalled),
		)

		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		router := mux.NewRouter()

		router.HandleFunc("/register", handler.handleRegister)
		router.ServeHTTP(rr, req)

		if rr.Code != http.StatusBadRequest {
			t.Errorf(
				"expected status code %d, got %d", 
				http.StatusBadRequest, rr.Code,
			) 
		}
	})

	t.Run("should register the user", func(t *testing.T) {
		payload := types.RegisterUserPayload {
			FirstName: "user",
			LastName: "smith",
			Email: "valid@mail.com",
			Password: "qwerty",
		}
		marchalled, _ := json.Marshal(payload)

		req, err := http.NewRequest(
			http.MethodPost, 
			"/register", 
			bytes.NewBuffer(marchalled),
		)

		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		router := mux.NewRouter()

		router.HandleFunc("/register", handler.handleRegister)
		router.ServeHTTP(rr, req)

		if rr.Code != http.StatusCreated {
			t.Errorf(
				"expected status code %d, got %d", 
				http.StatusBadRequest, rr.Code,
			) 
		}
	})
}

type mockUserStore struct {}

func (m *mockUserStore) GetUserByEmail(email string) (*types.User, error) {
	return nil, fmt.Errorf("no user found")
}

func (m *mockUserStore) GetUserByID(id int) (*types.User, error) {
	return nil, nil
}

func (m *mockUserStore) CreateUser(types.User) error {
	return nil
}