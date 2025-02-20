package models

import (
	"errors"
	"fmt"
)

type User struct {
	ID        int
	FirstName string
	LastName  string
}

var (
	users  []*User
	nextId = 1
)

func GetUsers() []*User {
	return users
}

func AddUser(user User) (User, error) {
	if user.ID != 0 {
		return User{}, errors.New("The new user must not contain any id")
	}
	user.ID = nextId
	nextId++

	users = append(users, &user)

	return user, nil
}

func GerUserById(id int) (User, error) {
	for _, user := range users {
		if user.ID == id {
			return *user, nil
		}
	}
	return User{}, fmt.Errorf("The user with id '%v' does not exist", id)
}

func UpdatedUserById(userToUpdate User) (User, error) {
	for i, user := range users {
		if user.ID == userToUpdate.ID {
			users[i] = &userToUpdate

			return userToUpdate, nil
		}
	}

	return User{}, fmt.Errorf("The user with id '%v' does not exist", userToUpdate.ID)
}

func RemoveUserByID(id int) ([]User, error) {
	for i, user := range users {
		if user.ID == id {
			users = append(users[:i], users[i+1:]...)

			return nil, nil
		}
	}

	return []User{}, fmt.Errorf("The user with id '%v' does not exist", id)
}
