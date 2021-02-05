package main

import (
	"github.com/chulista/subscription-addon/cmd"
)

func main() {
	application := cmd.NewApplication()
	application.Run(":8080")
}