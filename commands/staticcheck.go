package commands

import (
	"github.com/spf13/cobra"
)

var staticcheck = &cobra.Command{
	Use:   "staticcheck [files...]",
	Short: "Check your Go code with staticcheck",
	Long:  ``,
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		ensureInstalled("staticcheck", "honnef.co/go/tools/commands/staticcheck")
		runTool("staticcheck", pkgNames(dirNames(args)))
	},
}

func init() {
	rootCmd.AddCommand(staticcheck)
}
