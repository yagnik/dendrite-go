package generate

import (
	"../../command"
	"strings"
)

type GenerateConsulAgentCommand struct {
	command.Meta
}

func (c *GenerateConsulAgentCommand) Run(args []string) int {
	// Write your code here

	return 0
}

func (c *GenerateConsulAgentCommand) Synopsis() string {
	return "generate config for consul-agent"
}

func (c *GenerateConsulAgentCommand) Help() string {
	helpText := `

`
	return strings.TrimSpace(helpText)
}
