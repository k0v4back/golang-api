package services

import (
	"../entities"
	"../repositories"
	"crypto/sha1"
	"encoding/base64"
	"errors"
	"log"
	"time"
)

func GetAllUsers() ([]*entities.User, error) {
	return repositories.GetAllUsers()
}

func RegisterUser(Password, Email string) error {
	hash := sha1.New()

	CreatedAt := int(time.Now().Unix())
	ConfirmToken := base64.URLEncoding.EncodeToString(hash.Sum(nil))
	ConfirmTokenExpire := int(time.Now().Unix()) + 3600
	Status := entities.IS_WAITING

	if entities.CheckEmail(Email) != nil {
		return entities.CheckEmail(Email)
	} else if CheckUnicEmail(Email) != nil {
		return CheckUnicEmail(Email)
	}

	repositories.CreateUser(CreatedAt, ConfirmToken, ConfirmTokenExpire, Status, Password, Email)

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