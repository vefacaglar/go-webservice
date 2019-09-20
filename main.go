package main

import (
	"fmt"
	"net/http"

	"github.com/vefacaglar/webservice/controllers"
)

func main() {
	fmt.Println("started")
	controllers.RegisterControllers()
	http.ListenAndServe(":3000", nil)
}
