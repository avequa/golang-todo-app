package users_transport_http

import (
	"encoding/json"
	"fmt"
	"net/http"

	core_http_request "github.com/avequa/golang-todo-app/internal/core/transport/http/request"
)

type CreateUserRequest struct {
	FullName string `json:"full_name" validate:"required,min=3,max=100"`
	PhoneNumber *string `json:"phone_number" validate:"omitempty,min=3,max=100,startswith=+"`
}

type CreateUserResponse struct {
	ID int `json:"id"`
	Version int `json:"version"`
	FullName string `json:"full_name"`
	PhoneNumber *string `json:"phone_number"`
}

// POST /users
// request_id

func (h *UsersHTTPHandler) CreateUser(rw http.ResponseWriter, r *http.Request) {
	var request CreateUserRequest
	if err := core_http_request.DecodeAndValidateRequest(r, &request); err != nil {
		///
	}

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		fmt.Println("ошибка")
	}

	rw.WriteHeader(http.StatusOK)
}