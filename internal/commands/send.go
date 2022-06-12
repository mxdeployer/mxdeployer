package commands

import (
	"context"
	"fmt"
	"io/ioutil"

	"github.com/Azure/azure-sdk-for-go/sdk/messaging/azservicebus"
	"github.com/mxdeployer/mxdeployer/internal/core"
)

type Send struct {
	jsonPath string
}

func NewSend(argQueue *core.StringQueue) core.Command {
	return &Send{argQueue.Dequeue()}
}

func (cmd *Send) Run() error {

	config := core.LoadConfig()

	jsonBytes, err := ioutil.ReadFile(cmd.jsonPath)

	if err != nil {
		return err
	}
	client, err := azservicebus.NewClientFromConnectionString(config.AzServiceBusConStr, nil)

	if err != nil {
		return err
	}

	sender, err := client.NewSender(core.SbTopic, nil)

	if err != nil {
		return err
	}

	// TODO: move all of this into notification_queue

	err = sender.SendMessage(context.Background(), &azservicebus.Message{Body: jsonBytes, Subject: &config.Host}, nil)

	if err != nil {
		return err
	}

	fmt.Println("OK.")

	return nil
}
