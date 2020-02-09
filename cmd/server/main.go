package main

import (
	"log"

	"github.com/ivan-kapelyukh/knock-knock-unlock/internal/app/server"
	"github.com/ivan-kapelyukh/knock-knock-unlock/internal/app/server/mongo"
)

const serverPort = 8080
const staticDir = "./web/dist/"
const mongoHost = "localhost"
const mongoPort = 27017

func main() {
	mongo, err := mongo.New(mongoHost, mongoPort)

	if err != nil {
		log.Fatal(err)
	}

	server := server.New(serverPort, staticDir, mongo)
	log.Fatal(server.Serve())
}
