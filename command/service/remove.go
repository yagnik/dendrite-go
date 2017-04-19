package service

import (
	"../../command"
	"strings"
)

type ServiceRemoveCommand struct {
	command.Meta
}

func (c *ServiceRemoveCommand) Run(args []string) int {
	// Write your code here

	return 0
}

func (c *ServiceRemoveCommand) Synopsis() string {
	return "remove an existence service from file"
}

func (c *ServiceRemoveCommand) Help() string {
	helpText := `

`
	return strings.TrimSpace(helpText)
}
