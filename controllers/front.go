package controllers

import "net/http"

func RegisterControllers() {
	uc := NewUserController()

	http.Handle("/users", *uc)
	http.Handle("/users/", *uc)
}
