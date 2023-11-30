package adapters

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"apiGoHttp/pkg/core"
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
			log.Println(err.Error())
		}

		var bodyData core.LoginInfo

		err = json.Unmarshal(body, &bodyData)

		if err != nil {
			log.Panic(err)
		}

		routes := core.Routes{LoginInfo: bodyData}

		login, err := routes.Login()

		if err != nil {
			response.Status = true
			response.Message = err.Error()
		} else {
			response.Status = false
			response.Message = login
		}

		marshalResponse, err := json.Marshal(response)

		if err != nil {
			log.Panic("An error ocurred marshaling data in jason", err.Error())
		}

		res.Write(marshalResponse)
	} else {

		response.Status = false
		response.Message = "This api only acept a POST method"

		marshalResponse, err := json.Marshal(response)
		if err != nil {
			log.Panic("An error ocurred marshaling data in jason", err.Error())
		}

		res.Write(marshalResponse)
	}
}
