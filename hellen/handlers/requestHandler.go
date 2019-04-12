package handlers

import (
	"github.com/ASV44/Hellen-StandUp/hellen"
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