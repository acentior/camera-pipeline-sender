package main

import (
	"log"
	"os"

	vidoestreamsender "github.com/acentior/camera-pipeline-sender/internal/videoStreamSender"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Failed to load env {%v}", err)
		return
	}
	websocketUrl := os.Getenv("WEBSOCKET_URL")
	stunUrl := os.Getenv("STURN_URL")
	vss := vidoestreamsender.VideoStreamSender{}
	err := vss.Init(websocketUrl, stunUrl)
	if err != nil {
		log.Default().Fatalf("Failed to init: %v", err)
	}
	vss.Run()
}
