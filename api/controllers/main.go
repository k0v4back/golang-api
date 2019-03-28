package controllers

import (
	"../core/entities"
	"../core/helpers"
	"../core/repositories"
	"../core/services"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

var results []string

func MainPage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(helpers.SuccessResponse{Message: "Api 1.0 version"})
}

func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	value, err := repositories.GetAllUsers()
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}
	json.NewEncoder(w).Encode(helpers.UsersResponse{Users: value})
}

func UserRegister(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body",
			http.StatusInternalServerError)
	}
	results = append(results, string(body))

	data := []byte(results[0])
	u := &entities.User{}
	json.Unmarshal(data, u)

	if services.RegisterUser(u.Password, u.Email) != nil {
		json.NewEncoder(w).Encode(helpers.ErrorResponse{Error: "Invalid email address or this email already exist"})
	}else {
		json.NewEncoder(w).Encode(helpers.SuccessResponse{Message: "A confirmation email has been sent."})
	}
}

func ConfirmEmailByToken(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body",
			http.StatusInternalServerError)
	}
	results = append(results, string(body))

	data := []byte(results[0])
	u := &entities.User{}
	json.Unmarshal(data, u)

	e := services.CheckToken(u.ConfirmToken)
	//_, e := services.CheckToken(u.ConfirmToken)
	if e != nil {
		log.Fatal(err)
	}
}
