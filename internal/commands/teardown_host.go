package commands

import (
	"context"
	"fmt"

	"github.com/Azure/azure-sdk-for-go/sdk/messaging/azservicebus/admin"
	"github.com/mxdeployer/mxdeployer/internal/core"
)

type TeardownHost struct {
}

func NewTeardownHost() core.Command {
	return &TeardownHost{}
}

func (cmd *TeardownHost) Run() error {

	config := core.LoadConfig()

	const topic = "sbt-mxdeployer"
	sub := config.Host

	client, err := admin.NewClientFromConnectionString(config.AzServiceBusConStr, nil)

	if err != nil {
		return err
	}

	response, err := client.GetSubscription(context.Background(), topic, sub, nil)

	if err != nil {
		return err
	}

	if response != nil {

		fmt.Println("Deleting subscription for host...")

		_, err = client.DeleteSubscription(context.Background(), topic, sub, nil)

		if err != nil {
			return err
		}

		fmt.Println("OK.")
	}

	// TODO: remove service from system

	return nil
}
