package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/ASV44/Hellen-StandUp/hellen"
	"github.com/ASV44/Hellen-StandUp/hellen/slack/models"
	"io"
	"net/http"
	"time"
)

type RequestHandler struct {
	eventsHandlers map[string]func(w http.ResponseWriter, event models.Callback)
}

func NewRequestHandler() *RequestHandler {
	return &RequestHandler{eventsHandlers: initEventHandlers()}
}

func initEventHandlers() map[string]func(w http.ResponseWriter, event models.Callback) {
	eventsHandlers := make(map[string]func(w http.ResponseWriter, event models.Callback))
	eventsHandlers[models.UrlVerification] = confirmUrl
	eventsHandlers[models.EventCallback] = handleEvent
	return eventsHandlers
}

func (handler *RequestHandler) RootHandler(w http.ResponseWriter, r *http.Request) {
	currentTime := time.Now().Format("2006-01-02 15:04:05")
	_, _ = io.WriteString(w, hellen.WelcomeMessage+hellen.TimeMessage+currentTime)
}

func (handler *RequestHandler) SlackHandler(w http.ResponseWriter, r *http.Request) {
	var slackData models.Callback
	decodeError := json.NewDecoder(r.Body).Decode(&slackData)
	if decodeError != nil {
		fmt.Println(decodeError)
		return
	}
	data, _ := json.Marshal(slackData)
	fmt.Println(string(data))
	handler.eventsHandlers[slackData.Type](w, slackData)
}