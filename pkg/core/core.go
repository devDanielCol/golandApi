package core

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
)

type LoginInfo struct {
	Email    string `json:"name"`
	Password string `json:"password"`
}

type Routes struct {
	LoginInfo
}

const (
	password = "123456"
)

func (r *Routes) Login() (string, error) {
	if r.Password != password {
		return "", errors.New("invalid credentials")
	}

	return "Success Login", nil

}

func (r *Routes) UserInfo(res http.ResponseWriter, req http.Request) {

	userDefault := LoginInfo{
		Email:    "goland@google.com",
		Password: "12345678",
	}

	jsonUserData, err := json.Marshal(userDefault)

	if err != nil {
		log.Panic("Error parsing user data to json")
	}

	res.Write(jsonUserData)
}
