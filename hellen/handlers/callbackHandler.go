package handlers

import (
	"fmt"
	"github.com/ASV44/Hellen-StandUp/hellen/slack/models"
	"io"
	"net/http"
)

type CallbackHandler struct {
	eventsHandler *EventsHandler
	callbackHandlers map[string]func(w http.ResponseWriter, event models.Callback)
}

func NewCallbackHandler() *CallbackHandler {
	handler := &CallbackHandler{eventsHandler: NewEventsHandler()}
	handler.init()
	return handler
}

func (handler *CallbackHandler) init() {
	handler.callbackHandlers = make(map[string]func(w http.ResponseWriter, event models.Callback))
	handler.callbackHandlers[models.UrlVerification] = handler.confirmUrl
	handler.callbackHandlers[models.EventCallback] = handler.handleEvent
}

func (handler *CallbackHandler) ProcessCallback(w http.ResponseWriter, slackData models.Callback) {
	handler.callbackHandlers[slackData.Type](w, slackData)
}

func (handler *CallbackHandler) confirmUrl(w http.ResponseWriter, slackData models.Callback) {
	_, err := io.WriteString(w, slackData.Challenge)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func (handler *CallbackHandler) handleEvent(w http.ResponseWriter, slackData models.Callback) {
	handler.eventsHandler.ProcessEvent(slackData.Event)
}