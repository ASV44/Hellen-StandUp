package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/ASV44/Hellen-StandUp/hellen"
	"github.com/ASV44/Hellen-StandUp/hellen/slack"
	"github.com/ASV44/Hellen-StandUp/hellen/slack/models"
	"io"
	"net/http"
	"time"
)

type RequestHandler struct {

}

func NewRequestHandler() *RequestHandler {
	return &RequestHandler{}
}

func (handler *RequestHandler) RootHandler(w http.ResponseWriter, r *http.Request) {
	currentTime := time.Now().Format("2006-01-02 15:04:05")
	_, _ = io.WriteString(w, hellen.WelcomeMessage+hellen.TimeMessage+currentTime)
}

func (handler *RequestHandler) SlackHandler(w http.ResponseWriter, r *http.Request) {
	var slackData models.SlackCallback
	decodeError := json.NewDecoder(r.Body).Decode(&slackData)
	if decodeError != nil {
		fmt.Println(decodeError)
		return
	}
	data, _ := json.Marshal(slackData)
	fmt.Println(string(data))
	if slackData.Type == "url_verification" {
		_, _ = io.WriteString(w, slackData.Challenge)
		return
	}
	slack.SendMessage(slackData.Event.Channel, "Fucking Amazing")
}