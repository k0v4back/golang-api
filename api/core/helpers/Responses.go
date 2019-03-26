package helpers

import "../entities"

type ErrorResponse struct {
	Error 	string `json:"error"`
}

type SuccessResponse struct {
	Message string `json:"message"`
}

type UsersResponse struct {
	Users []*entities.User
}

type UsersCreate struct {
	Users []string
}