package main

import "github.com/flavio-foa-dev/adopet/server"

func main() {
	server := server.NewServer()
	server.Run()
}
