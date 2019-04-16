package handlers

import (
	"fmt"
	"github.com/ASV44/Hellen-StandUp/hellen/slack"
	"github.com/ASV44/Hellen-StandUp/hellen/slack/models"
	"io"
	"net/http"
)

func confirmUrl(w http.ResponseWriter, slackData models.Callback) {
	_, err := io.WriteString(w, slackData.Challenge)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func handleEvent(w http.ResponseWriter, slackData models.Callback)  {
	if slackData.Event.User != "" {
		slack.SendMessage(slackData.Event.Channel, "Amazing")
	}
}