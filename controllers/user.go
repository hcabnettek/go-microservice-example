package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/hcabnettek/filmapi/errors"
	"github.com/hcabnettek/filmapi/services"
)

type controller struct{}

var (
	userService services.UserService
)

// UserController interface
type UserController interface {
	GetUsers(response http.ResponseWriter, request *http.Request)
	AddUser(response http.ResponseWriter, request *http.Request)
}

// NewUserController constructor function returns UserController interface
func NewUserController(service services.UserService) UserController {
	userService = service
	return &controller{}
}

func (*controller) GetUsers(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	users, err := userService.FindAll()
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(response).Encode(errors.ServiceError{Message: "Error getting the users"})
	}
	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(users)
}

func (*controller) AddUser(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	/*posts, err := postService.FindAll()
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(response).Encode(errors.ServiceError{Message: "Error getting the posts"})
	}
	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(posts) */
}
