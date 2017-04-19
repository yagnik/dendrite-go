package service

import (
	"testing"

	"github.com/mitchellh/cli"
)

func TestServiceMetadataCommand_implement(t *testing.T) {
	var _ cli.Command = &ServiceMetadataCommand{}
}
