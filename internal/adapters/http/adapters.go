package httpAdapter

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	core "apiGoHttp/internal/core"
)

type Adapters struct{}

type ResponseData struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
}

var response *ResponseData

func (h *Adapters) Login(res http.ResponseWriter, req *http.Request) {

	if req.Method == "POST" {
		defer req.Body.Close()

		body, err := io.ReadAll(req.Body)

		if err != nil {
			http.Error(res, "Error in body reading", http.StatusInternalServerError)
			log.Panic("Not user data registered")
		}

		if len(body) == 0 {
			res.Write([]byte("Body cant be empty"))
			return
		}

		var bodyData core.LoginInfo

		err = json.Unmarshal(body, &bodyData)

		if err != nil {
			log.Panic(err)
		}

		routes := core.Routes{LoginInfo: bodyData}

		login, err := routes.Login()

		if err != nil {
			response = &ResponseData{Status: false, Message: "Error in login: " + err.Error()}
		} else {
			response = &ResponseData{Status: true, Message: login}
		}

		marshalResponse, err := json.Marshal(response)

		if err != nil {
			log.Panic("An error ocurred marshaling data in jason", err.Error())
		}

		res.Write(marshalResponse)
	} else {
		response = &ResponseData{Status: false, Message: "This api only acept a POST method"}
		marshalResponse, err := json.Marshal(response)
		if err != nil {
			log.Panic("An error ocurred marshaling data in jason", err.Error())
		}

		res.Write(marshalResponse)
	}
}
