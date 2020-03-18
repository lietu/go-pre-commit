package commands

import (
	"github.com/spf13/cobra"
)

var errcheck = &cobra.Command{
	Use:   "errcheck [files...]",
	Short: "Check your Go source code with errcheck",
	Long:  ``,
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		ensureInstalled("errcheck", "github.com/kisielk/errcheck")
		runTool("errcheck", pkgNames(dirNames(args)))
	},
}

func init() {
	rootCmd.AddCommand(errcheck)
}
