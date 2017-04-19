package main

import (
	"./command"
	"./command/generate"
	"./command/service"
	"github.com/mitchellh/cli"
)

func Commands(meta *command.Meta) map[string]cli.CommandFactory {
	return map[string]cli.CommandFactory{
		"service add": func() (cli.Command, error) {
			return &service.ServiceAddCommand{
				Meta: *meta,
			}, nil
		},
		"service remove": func() (cli.Command, error) {
			return &service.ServiceRemoveCommand{
				Meta: *meta,
			}, nil
		},
		"service this-depends-on": func() (cli.Command, error) {
			return &service.ServiceThisDependsOnCommand{
				Meta: *meta,
			}, nil
		},
		"service depends-on-this": func() (cli.Command, error) {
			return &service.ServiceDependsOnThisCommand{
				Meta: *meta,
			}, nil
		},
		"service metadata": func() (cli.Command, error) {
			return &service.ServiceMetadataCommand{
				Meta: *meta,
			}, nil
		},
		"generate nerve": func() (cli.Command, error) {
			return &generate.GenerateNerveCommand{
				Meta: *meta,
			}, nil
		},
		"generate synapse": func() (cli.Command, error) {
			return &generate.GenerateSynapseCommand{
				Meta: *meta,
			}, nil
		},
		"generate terraform": func() (cli.Command, error) {
			return &generate.GenerateTerraformCommand{
				Meta: *meta,
			}, nil
		},
		"generate consul-agent": func() (cli.Command, error) {
			return &generate.GenerateConsulAgentCommand{
				Meta: *meta,
			}, nil
		},
		"generate consul-template": func() (cli.Command, error) {
			return &generate.GenerateConsulTemplateCommand{
				Meta: *meta,
			}, nil
		},

		"version": func() (cli.Command, error) {
			return &command.VersionCommand{
				Meta:     *meta,
				Version:  Version,
				Revision: GitCommit,
				Name:     Name,
			}, nil
		},
	}
}
