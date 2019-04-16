package handlers

import (
	"github.com/ASV44/Hellen-StandUp/hellen/slack"
	"github.com/ASV44/Hellen-StandUp/hellen/slack/models"
)

type EventsHandler struct {
	eventsHandlers map[string]func(event models.Event)
}

func NewEventsHandler() *EventsHandler {
	handler := &EventsHandler{}
	handler.init()
	return handler
}

func (handler *EventsHandler) init() {
	handler.eventsHandlers = make(map[string]func(event models.Event))
	handler.eventsHandlers[models.Message] = handler.handleMessage
}


func (handler *EventsHandler) ProcessEvent(event models.Event) {
	handler.eventsHandlers[event.Type](event)
}

func (handler *EventsHandler) handleMessage(event models.Event)  {
	if event.User != "" {
		slack.SendMessage(event.Channel, "Amazing")
	}
}