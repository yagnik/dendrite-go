package service

import (
	"../../command"
	"strings"
)

type ServiceAddCommand struct {
	command.Meta
}

func (c *ServiceAddCommand) Run(args []string) int {
	// Write your code here

	return 0
}

func (c *ServiceAddCommand) Synopsis() string {
	return "add a new service to file"
}

func (c *ServiceAddCommand) Help() string {
	helpText := `

`
	return strings.TrimSpace(helpText)
}
