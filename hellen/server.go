package hellen

import (
	"fmt"
	"github.com/ASV44/Hellen-StandUp/hellen/slack"
	"github.com/gorilla/mux"
	"net/http"
	"os"
)

const (
	DefaultHost = "localhost"
	DefaultPort = "8080"
)

type Server struct {
	host     string
	port     string
	router   *mux.Router
}

func InitServer(host string, port string, router *mux.Router) Server {
	if host == "" {
		host = DefaultHost
	}
	if port == "" {
		port = DefaultPort
	}

	return Server{
		host:     host,
		port:     port,
		router:   router,
	}
}

func (server *Server) Start() {
	go server.run(server.router)
	slack.SendMessage(slack.OpenUserCommunication("UEPGE7U59"), "My message from server")
}

func (server *Server) run(router *mux.Router) {
	fmt.Println("Hellen is running on port : " + server.port)
	var err = http.ListenAndServe(":" + server.port, router)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
