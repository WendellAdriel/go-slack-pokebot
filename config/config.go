package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

type Token struct {
	Token string `json:"token"`
}

func GetToken() string {
	var token Token
	file, err := ioutil.ReadFile("./config/token.json")

	if err != nil {
		log.Fatal("File doesn't exist")
	}

	err = json.Unmarshal(file, &token)
	if err != nil {
		log.Fatal("Cannot parse token.json")
	}

	return token.Token
}
