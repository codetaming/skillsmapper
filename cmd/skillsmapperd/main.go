package main

import (
	"github.com/codetaming/skillsmapper/internal/api"
	"github.com/codetaming/skillsmapper/internal/persistence/local"
	"log"
	"net/http"
	"os"
)

var (
	serverPort = os.Getenv("SERVER_PORT")
)

func init() {
	if serverPort == "" {
		serverPort = "8080"
		log.Printf("$SERVER_PORT not set, defaulting to %s", serverPort)
	}
}

func main() {

	logger := log.New(os.Stdout, "skillsmapper ", log.LstdFlags|log.Lshortfile)

	logger.Printf("starting skillsmapper")

	logger.Printf("configuring data store")
	dataStore, err := local.NewInMemoryDataStore(logger)
	if err != nil {
		logger.Fatalf("failed to create data store: %v", err)
	}

	a := api.NewAPI(logger, dataStore)

	logger.Printf("server starting on port %s", serverPort)
	err = http.ListenAndServe(":"+serverPort, a.Router)
	if err != nil {
		logger.Fatalf("server failed to start: %v", err)
	}
}
