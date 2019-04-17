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
	callbackHandler *CallbackHandler
}

func NewRequestHandler() *RequestHandler {
	return &RequestHandler{callbackHandler: NewCallbackHandler()}
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
	handler.callbackHandler.ProcessCallback(w, slackData)
}

func (handler *RequestHandler) StandUpOpenDialogHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte{})
	err := r.ParseForm()
	if err != nil {
		fmt.Println(err)
		return
	}
	for key, value := range r.Form {
		fmt.Println(key, value)
	}
	openDialog := models.OpenDialog{TriggerId: r.Form["trigger_id"][0]}
	textarea := models.DialogElement{Type: models.Textarea,
									 Label: "What did you work on?",
									 Name: "worked On",
									 Hint: "What did you work on today?"}
	openDialog.Dialog = models.Dialog{CallbackId: "dgfdsdfds",
									  Title: "Stand Up Dialog",
									  SubmitLabel: "CMD",
									  NotifyOnCancel:true,
									  State: "Stand Up",
									  Elements: []models.DialogElement{textarea}}
	slack.OpenDialog(openDialog)
}