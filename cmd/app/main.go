package main

import (
	"log"

	"github.com/GenM4/go-ify/internal/config"
	"github.com/GenM4/go-ify/internal/server"
)

func main() {
	log.Print("Reading Configs")
	gcfg := config.GCFG{}
	if err := gcfg.Init(); err != nil {
		log.Print("Failed to initialize configs")
		panic(err)
	}

	log.Print("Building Server")
	server := server.NewServer(gcfg.ServerConfig)

	log.Print("Initializing Server")
	if err := server.Init(); err != nil {
		log.Print("Failed to initailize server")
		log.Print(err)
		panic(err)
	}
	log.Print("Server Initialized!")

	if err := server.Serve(); err != nil {
		log.Print("Failed to start server")
		panic(err)
	}
}
