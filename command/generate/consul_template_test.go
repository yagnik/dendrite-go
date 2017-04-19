package generate

import (
	"testing"

	"github.com/mitchellh/cli"
)

func TestGenerateConsulTemplateCommand_implement(t *testing.T) {
	var _ cli.Command = &GenerateConsulTemplateCommand{}
}
