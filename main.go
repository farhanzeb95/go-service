package main

import (
	"fmt"
	"go-web-service/models"
)

func main() {
	u := models.User{
		ID:        1,
		FirstName: "Farhan Zeb",
		LastName:  "Malik",
	}

	fmt.Println(u)

}
