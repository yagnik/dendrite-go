package service

import (
	"testing"

	"github.com/mitchellh/cli"
)

func TestServiceAddCommand_implement(t *testing.T) {
	var _ cli.Command = &ServiceAddCommand{}
}
