package core

import (
	"context"
	"encoding/json"
	"log"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/messaging/azservicebus"
	"github.com/mxdeployer/mxdeployer/internal/models"
)

type NotificationQueue struct {
	constr string
	host   string
	client *azservicebus.Client
}

type HandleNotificationFunc func(dn models.DeploymentNotification)

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

	if sender.SendMessage(context.Background(), &azservicebus.Message{Body: jsonBytes, Subject: &queue.host}, nil); err != nil {
		return err
	}

	return nil
}

// TODO: Need some way to handle errors better - backoff?
func (queue *NotificationQueue) Process(ctx context.Context, handler HandleNotificationFunc) {

	if err := queue.prepclient(); err != nil {
		log.Println(err)
	}

	rcvr, err := queue.client.NewReceiverForSubscription(SbTopic, queue.host, nil)

	if err != nil {
		log.Println(err)
	}

	for {
		msgs, err := rcvr.ReceiveMessages(ctx, 16, nil)

		if err != nil {
			if err == context.Canceled {
				return
			}
			log.Println(err)
		}

		for _, msg := range msgs {

			var dn models.DeploymentNotification
			if err := json.Unmarshal(msg.Body, &dn); err != nil {
				log.Println(err)
			}

			// Should this be executed as a goroutine?
			handler(dn)

			if err := rcvr.CompleteMessage(ctx, msg, nil); err != nil {
				if err == context.Canceled {
					return
				}
				log.Println(err)
			}
		}
	}
}

func (queue *NotificationQueue) Receive(timeout int) (*models.DeploymentNotification, error) {

	if err := queue.prepclient(); err != nil {
		return nil, err
	}

	rcvr, err := queue.client.NewReceiverForSubscription(SbTopic, queue.host, nil)

	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.TODO(), time.Duration(timeout)*time.Second)
	defer cancel()

	msgs, err := rcvr.ReceiveMessages(ctx, 1, nil)

	if err != nil {
		if err != context.DeadlineExceeded {
			return nil, err
		}
	}

	if len(msgs) > 0 {

		msg := msgs[0]

		var dn models.DeploymentNotification
		if err := json.Unmarshal(msg.Body, &dn); err != nil {
			return nil, err
		}

		if err := rcvr.CompleteMessage(context.TODO(), msg, nil); err != nil {
			return nil, err
		}

		return &dn, nil
	}

	return nil, nil
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
