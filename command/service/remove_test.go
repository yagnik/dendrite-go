package service

import (
	"testing"

	"github.com/mitchellh/cli"
)

func TestServiceRemoveCommand_implement(t *testing.T) {
	var _ cli.Command = &ServiceRemoveCommand{}
}
