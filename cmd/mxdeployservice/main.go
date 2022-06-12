package main

import (
	"github.com/mxdeployer/mxdeployer/internal/core"
)

func main() {
	queue := core.NewNotificationQueue("", "")
	queue.Receive()
}
