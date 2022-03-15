package main

import (
	"webapi-with-go/database"
	"webapi-with-go/server"
)

func main() {
	database.StartDatabase()

	server := server.NewServer()

	server.Run()
}
