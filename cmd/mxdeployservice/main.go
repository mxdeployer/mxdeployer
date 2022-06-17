package main

import (
	"log"
	"time"

	"github.com/mxdeployer/mxdeployer/internal/logfactory"
	"github.com/mxdeployer/mxdeployer/internal/service"
)

func main() {
	logwriter, err := logfactory.NewLogWriter()

	if err != nil {
		log.Fatal(err)
	}

	log.SetOutput(logwriter)

	service.Run(&mxdeployservice{})
}

type mxdeployservice struct {
	ticker *time.Ticker
}

func (s *mxdeployservice) Name() string {
	return "mxdeployservice"
}

func (s *mxdeployservice) Start() {

	log.Println("Starting...")

	s.ticker = time.NewTicker(1 * time.Second)
	go s.tick()
}

func (s *mxdeployservice) Stop() {

	log.Println("Stopping...")

	s.ticker.Stop()
}

func (s *mxdeployservice) tick() {

	for range s.ticker.C {
		log.Println("Tick!")
	}
}
