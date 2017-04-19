package generate

import (
	"../../command"
	"strings"
)

type GenerateSynapseCommand struct {
	command.Meta
}

func (c *GenerateSynapseCommand) Run(args []string) int {
	// Write your code here

	return 0
}

func (c *GenerateSynapseCommand) Synopsis() string {
	return "generate config for smartstack synapse"
}

func (c *GenerateSynapseCommand) Help() string {
	helpText := `

`
	return strings.TrimSpace(helpText)
}
