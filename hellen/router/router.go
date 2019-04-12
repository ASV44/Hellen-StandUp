package router

import (
	"github.com/ASV44/Hellen-StandUp/hellen/handlers"
	"github.com/gorilla/mux"
)

func New(handler *handlers.RequestHandler) *mux.Router {
	router := mux.NewRouter()
	addRoutesHandlers(router, handler)

	return router
}

func addRoutesHandlers(router *mux.Router, handler *handlers.RequestHandler) {
	router.HandleFunc("/", handler.RootHandler)
}