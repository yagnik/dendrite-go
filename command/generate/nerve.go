package generate

import (
	"../../command"
	"strings"
)

type GenerateNerveCommand struct {
	command.Meta
}

func (c *GenerateNerveCommand) Run(args []string) int {
	// Write your code here

	return 0
}

func (c *GenerateNerveCommand) Synopsis() string {
	return "generate config for smartstack nerve"
}

func (c *GenerateNerveCommand) Help() string {
	helpText := `

`
	return strings.TrimSpace(helpText)
}
