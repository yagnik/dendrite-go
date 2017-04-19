package service

import (
	"../../command"
	"strings"
)

type ServiceMetadataCommand struct {
	command.Meta
}

func (c *ServiceMetadataCommand) Run(args []string) int {
	// Write your code here

	return 0
}

func (c *ServiceMetadataCommand) Synopsis() string {
	return "output all data related to the service"
}

func (c *ServiceMetadataCommand) Help() string {
	helpText := `

`
	return strings.TrimSpace(helpText)
}
