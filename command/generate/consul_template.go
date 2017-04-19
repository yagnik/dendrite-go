package generate

import (
	"../../command"
	"strings"
)

type GenerateConsulTemplateCommand struct {
	command.Meta
}

func (c *GenerateConsulTemplateCommand) Run(args []string) int {
	// Write your code here

	return 0
}

func (c *GenerateConsulTemplateCommand) Synopsis() string {
	return "generate config for consul-template"
}

func (c *GenerateConsulTemplateCommand) Help() string {
	helpText := `

`
	return strings.TrimSpace(helpText)
}
