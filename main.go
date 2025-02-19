package main

import (
	"cryptoScamDetection/config"
	"cryptoScamDetection/server"
)

func main() {
	config.LoadEnv()
	server.InitServer()

}
