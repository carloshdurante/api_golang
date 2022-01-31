package main

import (
	"github.com/carloshdurante/api_golang/database"
	"github.com/carloshdurante/api_golang/server"
)

func main() {
	database.StartDB()

	server := server.NewServer()

	server.Run()
}
