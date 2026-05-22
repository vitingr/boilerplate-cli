package cmd

import (
	"os"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "bplt",
	Short: "Boilerplate CLI - scaffold projects interactively",
	Long: color.CyanString(`
		██████╗ ██████╗ ██╗  ████████╗
		██╔══██╗██╔══██╗██║  ╚══██╔══╝
		██████╔╝██████╔╝██║     ██║   
		██╔══██╗██╔═══╝ ██║     ██║   
		██████╔╝██║     ███████╗██║   
		╚═════╝ ╚═╝     ╚══════╝╚═╝   
	`) + "\nInteractive project scaffolding from your boilerplate repo.",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(createCmd)
}
