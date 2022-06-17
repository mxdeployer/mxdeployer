package main

import (
	"log"

	"github.com/mxdeployer/mxdeployer/internal/logfactory"
)

func main() {
	logwriter, err := logfactory.NewLogWriter()

	if err != nil {
		log.Fatal(err)
	}

	log.SetOutput(logwriter)

	log.Println("Shutting down!")
}
