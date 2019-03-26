package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"./controllers"
)

func main()  {
	r := mux.NewRouter()

	r.HandleFunc("/", controllers.MainPage).Methods("GET")
	r.HandleFunc("/all-users", controllers.GetAllUsers).Methods("GET")
	r.HandleFunc("/user-register", controllers.UserRegister).Methods("POST")

	fmt.Println("Listen port 8080")

	fmt.Println()

	log.Fatal(http.ListenAndServe(":8080", r))
}