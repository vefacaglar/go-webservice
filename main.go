package main

import (
	"fmt"

	"github.com/vefacaglar/webservice/models"
)

func main() {
	fmt.Println("running...")
	user := models.User{
		ID:        2,
		FirstName: "Vefa",
		LastName:  "Caglar",
	}

	fmt.Println(user)
}
