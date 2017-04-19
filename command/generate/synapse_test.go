package generate

import (
	"testing"

	"github.com/mitchellh/cli"
)

func TestGenerateSynapseCommand_implement(t *testing.T) {
	var _ cli.Command = &GenerateSynapseCommand{}
}
