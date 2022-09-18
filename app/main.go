package main

import (
	"apiserver/server"
	"apiserver/server/logger"
)

func main() {
	server := server.NewServer(8080, logger.Debug)

	server.Start()
}
