package generate

import (
	"../../command"
	"strings"
)

type GenerateTerraformCommand struct {
	command.Meta
}

func (c *GenerateTerraformCommand) Run(args []string) int {
	// Write your code here

	return 0
}

func (c *GenerateTerraformCommand) Synopsis() string {
	return "generate config for terraform"
}

func (c *GenerateTerraformCommand) Help() string {
	helpText := `

`
	return strings.TrimSpace(helpText)
}
