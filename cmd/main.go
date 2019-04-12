package main

import (
	"github.com/ASV44/Hellen-StandUp/hellen"
	"github.com/ASV44/Hellen-StandUp/hellen/handlers"
	"github.com/ASV44/Hellen-StandUp/hellen/router"
	"os"
)

func main() {
	requestHandler := handlers.NewRequestHandler()
	serverRouter := router.New(requestHandler)
	instance := hellen.InitServer("", os.Getenv("PORT"), serverRouter)
	instance.Start()
}