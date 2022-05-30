package main

import (
	"fmt"
	"os"

	"github.com/mxdeployer/mxdeployer/internal/core"
)

func showHelp() {
	fmt.Println("deployctrl v0.1")
	fmt.Println()
	fmt.Println("Usage: deployctrl <command> [parameters]")
	fmt.Println()
	fmt.Println("Commands:")
	fmt.Println()
	fmt.Println("  setup")
	fmt.Println("      Create underlying storage entities.")
	fmt.Println()
	fmt.Println("  send <json-filename>")
	fmt.Println("      Send an artifact-ready notification.")
	fmt.Println()
	fmt.Println("  receive")
	fmt.Println("      Wait for a single artifact-ready notification and show it.")
	fmt.Println()
	fmt.Println("  drain")
	fmt.Println("      Drain the artifact queue.")
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

	args := os.Args[1:]
	var cmd core.Command

	switch cmdArg := args[0]; cmdArg {
	case "setup":
		cmd = core.NewSetupCommand()
	default:
		println("Unknown command.")
		println()
		showHelp()
		return
	}

	cmd.Run()
}
