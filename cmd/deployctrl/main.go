package main

import (
	"fmt"
	"os"

	"github.com/mxdeployer/mxdeployer/internal/commands"
	"github.com/mxdeployer/mxdeployer/internal/core"
)

func showHelp() {
	fmt.Println("deployctrl v0.1")
	fmt.Println()
	fmt.Println("Usage: deployctrl <command> [parameters]")
	fmt.Println()
	fmt.Println("Commands:")
	fmt.Println()
	fmt.Println("  setup-host <org> <az-storage-con-str> <az-service-bus-con-str> [host]")
	fmt.Println("      Setup a new deployment host. Creates a new artifact subscription.")
	fmt.Println()
	fmt.Println("  teardown-host")
	fmt.Println("      Deletes any existing artifact subscriptions.")
	fmt.Println()
	fmt.Println("  init")
	fmt.Println("      Initialize an empty deployment notification JSON document.")
	fmt.Println()
	fmt.Println("  send <json-filename>")
	fmt.Println("      Send an deployment notification.")
	fmt.Println()
	fmt.Println("  receive [host]")
	fmt.Println("      Waits for a single deployment notification.")
	fmt.Println()
	fmt.Println("  drain [host]")
	fmt.Println("      Drains an artifact subscription of all deployment notifications.")
	fmt.Println()
	fmt.Println("  purge")
	fmt.Println("      Purge all artifacts from blob storage.")
	fmt.Println()
	fmt.Println("  deploy <zip-filename> <json-filename>")
	fmt.Println("      Deploy a zip file using meta data specified in a json file.")
	fmt.Println()
}

func main() {

	if len(os.Args) < 2 {
		showHelp()
		return
	}

	argQueue := core.NewStringQueue(os.Args[1:])

	var cmd core.Command

	switch cmdArg := argQueue.Dequeue(); cmdArg {

	case "setup-host":
		cmd = commands.NewSetupHost(argQueue)
	case "teardown-host":
		cmd = commands.NewTeardownHost()
	case "init":
		cmd = commands.NewInit()
	case "send":
		cmd = commands.NewSend(argQueue)

	default:
		println("Unknown command.")
		println()
		showHelp()
		return
	}

	err := cmd.Run()

	if err != nil {
		fmt.Println(err)
	}
}
