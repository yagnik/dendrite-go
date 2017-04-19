package service

import (
	"../../command"
	"strings"
)

type ServiceDependsOnThisCommand struct {
	command.Meta
}

func (c *ServiceDependsOnThisCommand) Run(args []string) int {
	// Write your code here

	return 0
}

func (c *ServiceDependsOnThisCommand) Synopsis() string {
	return "list all services that depend on the service"
}

func (c *ServiceDependsOnThisCommand) Help() string {
	helpText := `

`
	return strings.TrimSpace(helpText)
}
