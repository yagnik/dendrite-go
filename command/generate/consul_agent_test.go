package generate

import (
	"testing"

	"github.com/mitchellh/cli"
)

func TestGenerateConsulAgentCommand_implement(t *testing.T) {
	var _ cli.Command = &GenerateConsulAgentCommand{}
}
