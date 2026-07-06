package users_transport_http

import (
	"fmt"
	"net/http"
	"encoding/json"
)

type CreateUserRequest struct {
	FullName string `json:"full_name"`
	PhoneNumber *string `json:"phone_number"`
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
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		fmt.Println("ошибка")
	}
}