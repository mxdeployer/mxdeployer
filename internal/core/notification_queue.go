package core

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/Azure/azure-sdk-for-go/sdk/messaging/azservicebus"
	"github.com/mxdeployer/mxdeployer/internal/models"
)

type NotificationQueue struct {
	constr string
	host   string
	client *azservicebus.Client
}

func NewNotificationQueue(constr string, host string) *NotificationQueue {
	return &NotificationQueue{constr, host, nil}
}

func (queue *NotificationQueue) Send(notification models.DeploymentNotification) error {

	err := queue.prepclient()

	if err != nil {
		return err
	}

	sender, err := queue.client.NewSender(SbTopic, nil)

	if err != nil {
		return err
	}

	jsonBytes, err := json.Marshal(notification)

	if err != nil {
		return err
	}

	err = sender.SendMessage(context.Background(), &azservicebus.Message{Body: jsonBytes, Subject: &queue.host}, nil)

	if err != nil {
		return err
	}

	return nil
}

func (queue *NotificationQueue) Receive() {
	fmt.Println("Receiving...")
}

func (queue *NotificationQueue) prepclient() error {
	if queue.client == nil {
		client, err := azservicebus.NewClientFromConnectionString(queue.constr, nil)

		if err != nil {
			return err
		}

		queue.client = client
	}

	return nil
}
