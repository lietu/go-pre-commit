package commands

import (
	"github.com/spf13/cobra"
)

var gofumports = &cobra.Command{
	Use:   "gofumports [files...]",
	Short: "Format your Go code with gofumports",
	Long:  ``,
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		ensureInstalled("gofumports", "mvdan.cc/gofumpt/gofumports")
		runTool("gofumports", append([]string{"-l", "-w"}, args...))
	},
}

func init() {
	rootCmd.AddCommand(gofumports)
}
