package main

import (
	"log"
	"schedulizer/server"
)

func main() {
	log.Fatalln(server.Run(":8090"))
}
