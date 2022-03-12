package main

import "webapi-with-go/server"

func main() {
	server := server.NewServer()

	server.Run()
}
