package entities

import (
	"errors"
	"regexp"
)

const ACTIVE = 1
const IS_WAITING = 2
const BLOCKED = 3

type User struct {
	Id                 int    `json:"id"`
	Username           string `json:"username"`
	Nick               string `json:"nick"`
	CreatedAt          int    `json:"created_at"`
	ConfirmToken       string `json:"confirm_token"`
	ConfirmTokenExpire int    `json:"confirm_token_expire"`
	Status             int    `json:"status"`
	Password           string `json:"password"`
	Email              string `json:"email"`
}

func CheckEmail(Email string) (error) {
	re := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	if re.MatchString(Email) == false {
		return errors.New("invalid email address or this email already exist")
	}

	return nil
}