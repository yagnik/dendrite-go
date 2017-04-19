package service

import (
	"../../command"
	"strings"
)

type ServiceThisDependsOnCommand struct {
	command.Meta
}

func (c *ServiceThisDependsOnCommand) Run(args []string) int {
	// Write your code here

	return 0
}

func (c *ServiceThisDependsOnCommand) Synopsis() string {
	return "list all services this service depends on"
}

func (c *ServiceThisDependsOnCommand) Help() string {
	helpText := `

`
	return strings.TrimSpace(helpText)
}
