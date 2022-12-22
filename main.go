package main

import "log"

func main() {
	server := InitServer()
	log.Fatal(server.Run(":3030"))
}
