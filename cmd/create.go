package cmd

import (
	"github.com/spf13/cobra"
	"github.com/vitingr/boilerplate-cli/internal/prompt"
	"github.com/vitingr/boilerplate-cli/internal/scaffold"
)

var createCmd = &cobra.Command{
	Use:   "create [project-name]",
	Short: "Create a new project from a boilerplate",
	Args:  cobra.MaximumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		projectName := ""
		if len(args) > 0 {
			projectName = args[0]
		}

		cfg, err := prompt.Run(projectName)
		if err != nil {
			return err
		}

		return scaffold.Run(cfg)
	},
}