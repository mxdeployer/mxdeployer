package core

import (
	"fmt"

	"github.com/mxdeployer/mxdeployer/internal/models"
)

type NotificationQueue struct {
	constr string
}

func NewNotificationQueue(constr string) *NotificationQueue {
	return &NotificationQueue{constr}
}

func (queue NotificationQueue) Send(notification *models.DeploymentNotification) {
	fmt.Printf("Sending '%s' ...\n", notification.Url)
}

func (queue NotificationQueue) Receive() {
	fmt.Println("Receiving...")
}
