package generate

import (
	"testing"

	"github.com/mitchellh/cli"
)

func TestGenerateTerraformCommand_implement(t *testing.T) {
	var _ cli.Command = &GenerateTerraformCommand{}
}
