package commands

import (
	"context"
	"fmt"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/messaging/azservicebus/admin"
	"github.com/mxdeployer/mxdeployer/internal/core"
)

type SetupHost struct {
	config core.Config
}

func NewSetupHost(argQueue *core.StringQueue) core.Command {

	cfg := core.LoadConfig()

	cfg.Org = argQueue.DequeueOrDefault(cfg.Org)
	cfg.AzStorageConStr = argQueue.DequeueOrDefault(cfg.AzStorageConStr)
	cfg.AzServiceBusConStr = argQueue.DequeueOrDefault(cfg.AzServiceBusConStr)
	cfg.Host = argQueue.DequeueOrDefault(cfg.Host)

	return &SetupHost{cfg}
}

func (cmd *SetupHost) Run() error {

	err := core.SaveConfig(cmd.config)

	if err != nil {
		return err
	}

	fmt.Printf("Configuration saved to: %s\n", core.ConfigPath())

	client, err := admin.NewClientFromConnectionString(cmd.config.AzServiceBusConStr, nil)

	if err != nil {
		return err
	}

	const topic = "sbt-mxdeployer"
	sub := cmd.config.Host

	response, _ := client.GetSubscription(context.Background(), topic, sub, nil)

	if response == nil {

		fmt.Println("Creating new subscription...")

		subOptions := &admin.CreateSubscriptionOptions{
			Properties: &admin.SubscriptionProperties{
				DefaultMessageTimeToLive: to.Ptr("P14D"),
			}}

		_, err = client.CreateSubscription(context.Background(), topic, sub, subOptions)

		if err != nil {
			return err
		}

		fmt.Println("OK.")
		fmt.Println("Creating correlation filter for host...")

		ruleOptions := &admin.CreateRuleOptions{
			Name: to.Ptr("match_host"),
			Filter: &admin.CorrelationFilter{
				Subject: &sub,
			},
		}

		_, err := client.CreateRule(context.Background(), topic, sub, ruleOptions)

		if err != nil {
			return err
		}

		fmt.Println("OK.")

		// TODO: register service with system
	}

	return nil
}
