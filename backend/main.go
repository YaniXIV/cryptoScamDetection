package main

import (
	"cryptoScamDetection/backend/config"
	"cryptoScamDetection/backend/server"
)

func main() {
	config.LoadEnv()
	server.InitServer()

}
