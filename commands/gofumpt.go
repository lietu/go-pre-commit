package commands

import (
	"github.com/spf13/cobra"
)

var gofumpt = &cobra.Command{
	Use:   "gofumpt [files...]",
	Short: "Format your Go code with gofumpt",
	Long:  ``,
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		ensureInstalled("gofumpt", "mvdan.cc/gofumpt")
		runTool("gofumpt", append([]string{"-l", "-w"}, args...))
	},
}

func init() {
	rootCmd.AddCommand(gofumpt)
}
