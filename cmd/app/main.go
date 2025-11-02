package main

import (
	//"fmt"

	//"github.com/GenM4/go-ify/internal/api"
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
	server := server.NewServer(&gcfg)

	log.Print("Initializing Server")
	server.Init()
	log.Print("Server Initialized!")

	if err := server.Serve(); err != nil {
		log.Print("Failed to start server")
		panic(err)
	}

	//--------------------------------  \/ figure out where this shit goes
	/*
		spotifyApi := api.SpotifyApiInit(gcfg)

		fmt.Println("Enter a spotify share URL:")
		var rawURL string
		_, err := fmt.Scanln(&rawURL)
		if err != nil {
			log.Fatal("Invalid input")
		}
		fmt.Println()

		asset := spotifyApi.ParseInput(rawURL)

		r := spotifyApi.GetSpotifyAsset(asset)

		r.Log()
	*/

}
