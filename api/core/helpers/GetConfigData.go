package helpers

import (
	"encoding/json"
	"fmt"
	"os"
)

type Configuration struct {
	Email    	string
	Password   	string
}

func getConfFile () Configuration {
	file, _ := os.Open("conf.json")
	decoder := json.NewDecoder(file)
	configuration := Configuration{}
	err := decoder.Decode(&configuration)
	if err != nil {
		fmt.Println("error:", err)
	}

	return configuration
}

func GetEmail() string {
	return getConfFile().Email
}

func GetEmailPassword() string {
	return getConfFile().Password
}