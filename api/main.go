package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"./controllers"
	"../api/core/helpers"
)

func main()  {
	r := mux.NewRouter()

	r.HandleFunc("/", helpers.JsonMiddleware(controllers.MainPage)).Methods("GET")
	r.HandleFunc("/all-users", helpers.JsonMiddleware(controllers.GetAllUsers)).Methods("GET")
	r.HandleFunc("/user-register", helpers.JsonMiddleware(controllers.UserRegister)).Methods("POST")
	r.HandleFunc("/confirm-email", helpers.JsonMiddleware(controllers.ConfirmEmailByToken)).Methods("POST")

	fmt.Println("Listen port 8080")

	fmt.Println()

	log.Fatal(http.ListenAndServe(":8080", r))
}