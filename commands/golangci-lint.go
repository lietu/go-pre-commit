package commands

import (
	"github.com/spf13/cobra"
)

var golangciLint = &cobra.Command{
	Use:   "golangci-lint [files...]",
	Short: "Check your Go code with golangci-lint",
	Long:  ``,
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		runTool("golangci-lint", append([]string{"run"}, pkgNames(dirNames(args))...))
	},
}

func init() {
	rootCmd.AddCommand(golangciLint)
}
