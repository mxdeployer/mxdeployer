package main

import (
	"context"
	"log"

	"github.com/mxdeployer/mxdeployer/internal/core"
	"github.com/mxdeployer/mxdeployer/internal/logfactory"
	"github.com/mxdeployer/mxdeployer/internal/models"
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
	queue  core.NotificationQueue
	cancel context.CancelFunc
}

func (s *mxdeployservice) Name() string {
	return "mxdeployservice"
}

func (s *mxdeployservice) Start() {

	log.Println("Starting...")

	cfg := core.LoadConfig()

	s.queue = *core.NewNotificationQueue(cfg.AzServiceBusConStr, cfg.Host)

	ctx := context.Background()
	var cancelCtx context.Context
	cancelCtx, s.cancel = context.WithCancel(ctx)

	go s.queue.Process(cancelCtx, s.NotificationReceived)
}

func (s *mxdeployservice) Stop() {

	log.Println("Stopping...")

	s.cancel()
}

func (s *mxdeployservice) NotificationReceived(dn models.DeploymentNotification) {
	log.Println("We got one!")
}
