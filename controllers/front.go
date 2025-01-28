package controllers

import "net/http"

func RegisterControllers() {
	userController := newUserController()

	http.Handle("/users", *userController)
	http.Handle("/users/", *userController)

}
