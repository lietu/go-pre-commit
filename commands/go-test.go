package commands

import (
	"github.com/spf13/cobra"
)

var goTest = &cobra.Command{
	Use:   "go-test [files...]",
	Short: "Run your Go unit tests",
	Long:  ``,
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		runTool("go", append([]string{"test"}, pkgNames(dirNames(args))...))
	},
}

func init() {
	rootCmd.AddCommand(goTest)
}
