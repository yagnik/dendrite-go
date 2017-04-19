package generate

import (
	"testing"

	"github.com/mitchellh/cli"
)

func TestGenerateNerveCommand_implement(t *testing.T) {
	var _ cli.Command = &GenerateNerveCommand{}
}
