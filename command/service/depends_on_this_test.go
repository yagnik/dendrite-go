package service

import (
	"testing"

	"github.com/mitchellh/cli"
)

func TestServiceDependsOnThisCommand_implement(t *testing.T) {
	var _ cli.Command = &ServiceDependsOnThisCommand{}
}
