package controllers

import (
	"encoding/json"
	"fmt"
	"go-web-service/models"
	"io"
	"net/http"
	"regexp"
	"strconv"
)

type userController struct {
	userIdPattern *regexp.Regexp
}

func (userController userController) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/users" {
		switch r.Method {
		case http.MethodGet:
			userController.getAllUsers(w, r)
		case http.MethodPost:
			userController.AddNewUser(w, r)
		default:
			w.WriteHeader(http.StatusNotImplemented)
		}
	} else {
		matches := userController.userIdPattern.FindStringSubmatch(r.URL.Path)
		fmt.Println(matches)

		if len(matches) < 2 { // Ensure at least two elements before accessing matches[1]
			w.WriteHeader(http.StatusNotFound)
			return
		}

		id, err := strconv.Atoi(matches[1])
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		switch r.Method {
		case http.MethodGet:
			userController.getUserById(id, w)
		case http.MethodPut:
			userController.UpdateUser(id, w, r)
		case http.MethodDelete:
			panic("method not implemented")
		default:
			w.WriteHeader(http.StatusNotImplemented)
		}
	}
}

func newUserController() *userController {
	return &userController{
		userIdPattern: regexp.MustCompile(`^/users/(\d+)/?$`), // Corrected regex
	}
}

func (userController userController) getAllUsers(w http.ResponseWriter, r *http.Request) {
	encodeResponseAsJSON(models.GetUsers(), w)
}

func encodeResponseAsJSON(data interface{}, w io.Writer) {
	encodedJson := json.NewEncoder(w)

	encodedJson.Encode((data))
}

func (userController userController) getUserById(id int, w http.ResponseWriter) {
	fmt.Println(id)
	user, error := models.GerUserById(id)

	if error != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	encodeResponseAsJSON(user, w)
}

func (userController userController) AddNewUser(w http.ResponseWriter, r *http.Request) {
	user, error := parseRequest(r)

	if error != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Could not pasrse the object"))
	}

	user, error = models.AddUser(user)

	if error != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(error.Error()))

		return
	}

	encodeResponseAsJSON(user, w)

}

func parseRequest(r *http.Request) (models.User, error) {
	decoder := json.NewDecoder((r.Body))
	var u models.User

	error := decoder.Decode(&u)

	if error != nil {
		return models.User{}, error
	}

	return u, nil

}

func (userController userController) UpdateUser(id int, w http.ResponseWriter, r *http.Request) {
	panic("method not implemented")
}

func (userController userController) RemoveUser(w http.ResponseWriter, r *http.Request) {
	panic("method not implemented")
}

// func newUserController() *userController {
// 	return &userController{
// 		userIdPattern: regexp.MustCompile(`^/users/(/d)/?`),
// 	}
// }
