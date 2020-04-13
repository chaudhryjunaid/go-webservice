package main

import (
	"fmt"

	"github.com/chaudhryjunaid/go-webservice/models"
)

func main() {
	u := models.User{
		ID:        2,
		FirstName: "Chaudhry",
		LastName:  "Anwar",
	}
	fmt.Println(u)
}
