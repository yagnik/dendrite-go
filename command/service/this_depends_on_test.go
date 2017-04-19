package service

import (
	"testing"

	"github.com/mitchellh/cli"
)

func TestServiceThisDependsOnCommand_implement(t *testing.T) {
	var _ cli.Command = &ServiceThisDependsOnCommand{}
}
