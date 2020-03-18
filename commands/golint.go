package commands

import (
	"github.com/spf13/cobra"
)

var golint = &cobra.Command{
	Use:   "golint [files...]",
	Short: "Check your Go code with golint",
	Long:  ``,
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		ensureInstalled("golint", "golang.org/x/lint/golint")
		runTool("golint", append([]string{"-set_exit_status"}, args...))
	},
}

func init() {
	rootCmd.AddCommand(golint)
}
