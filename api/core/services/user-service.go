package services

import (
	"../entities"
	"../helpers"
	"../repositories"
	"errors"
	"log"
	"time"
)

func GetAllUsers() ([]*entities.User, error) {
	return repositories.GetAllUsers()
}

func RegisterUser(Password, Email string) error {

	CreatedAt := int(time.Now().Unix())
	ConfirmToken := helpers.RandStringRunes(6)
	ConfirmTokenExpire := int(time.Now().Unix()) + 3600
	Status := entities.IS_WAITING

	if entities.CheckEmail(Email) != nil {
		return entities.CheckEmail(Email)
	} else if CheckUnicEmail(Email) != nil {
		return CheckUnicEmail(Email)
	}

	repositories.CreateUser(CreatedAt, ConfirmToken, ConfirmTokenExpire, Status, Password, Email)

	FromEmail := helpers.GetEmail()
	FromPassword := helpers.GetEmailPassword()
	ToEmail := Email
	Subject := "Confirm email by token"
	Body := "Confirm token: " + ConfirmToken
	helpers.SendEmail(FromEmail, FromPassword, ToEmail, Subject, Body)

	return nil
}

func CheckUnicEmail (Email string) error {
	value, err := repositories.GetAllUsers()
	if err != nil {
		log.Fatal(err)
	}

	for _, v := range value {
		if v.Email == Email {
			return errors.New("this email already exist")
		}
	}
	return nil
}


func CheckToken (Token string) error {
	value, err := repositories.GetAllUsers()
	if err != nil {
		log.Fatal(err)
	}

	for _, v := range value {
		if v.ConfirmToken == Token && v.ConfirmTokenExpire > int(time.Now().Unix()) {
			Status := entities.ACTIVE
			_, err := repositories.UpdateUserById(v.Id, Status, v.Username, v.Nick)
			if err != nil {
				return err
			}
			return nil
		}
	}
	return nil
}