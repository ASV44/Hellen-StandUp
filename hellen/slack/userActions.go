package slack

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/ASV44/Hellen-StandUp/hellen/slack/models"
	"net/http"
	"net/url"
	"os"
)

func OpenUserCommunication(userID string) string {
	urlString := API + "/im.open?" + "token=" + os.Getenv("BOT_TOKEN") + "&user=" + userID
	forwardReq, _ := http.NewRequest(http.MethodGet, urlString, nil)
	response, err := http.DefaultClient.Do(forwardReq)

	if err != nil {
		fmt.Print(err)
		return ""
	}

	var openChannelResponse models.OpenChannel
	_ = json.NewDecoder(response.Body).Decode(&openChannelResponse)
	id := openChannelResponse.Id()
	fmt.Println(id)
	return id
}

func SendMessage(channelID string, message string) {
	urlString := API + "/chat.postMessage?" + "token=" + os.Getenv("BOT_TOKEN") + "&channel=" + channelID + "&text=" + url.QueryEscape(message)
	req, _ := http.NewRequest(http.MethodPost, urlString, nil)
	_, _ = http.DefaultClient.Do(req)
}

func OpenDialog(openDialog models.OpenDialog) {
	urlString := API + "/dialog.open"
	data, _ := json.Marshal(openDialog)
	req, _ := http.NewRequest(http.MethodPost, urlString, bytes.NewReader(data))
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer " + os.Getenv("BOT_TOKEN"))
	_, _ = http.DefaultClient.Do(req)
}