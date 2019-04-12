package slack

import (
	"encoding/json"
	"fmt"
	"github.com/ASV44/Hellen-StandUp/hellen/slack/models"
	"net/http"
	"net/url"
)

func OpenUserCommunication(userID string) string {
	urlString := API + "/im.open?" + "token=" + BotToken + "&user=" + userID
	forwardReq, _ := http.NewRequest(http.MethodGet, urlString, nil)
	response, err := http.DefaultClient.Do(forwardReq)

	if err != nil {
		fmt.Print(err)
		return ""
	}

	var openChannelResponse models.OpenChannelResponse
	_ = json.NewDecoder(response.Body).Decode(&openChannelResponse)
	return openChannelResponse.Id()
}

func SendMessage(channelID string, message string) {
	urlString := API + "/chat.postMessage?" + "token=" + BotToken + "&channel=" + channelID + "&text=" + url.QueryEscape(message)
	forwardReq, _ := http.NewRequest(http.MethodPost, urlString, nil)
	_, _ = http.DefaultClient.Do(forwardReq)
}