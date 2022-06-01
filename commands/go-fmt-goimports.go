package commands

import (
	"github.com/spf13/cobra"
)

var goFmtGoimports = &cobra.Command{
	Use:   "go-fmt-goimports [files...]",
	Short: "Format your Go code with go fmt and goimports",
	Long:  ``,
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		ensureInstalled("goimports", "golang.org/x/tools/cmd/goimports")
		runTool("goimports", append([]string{"-l", "-w"}, args...))
	},
}

func init() {
	rootCmd.AddCommand(goFmtGoimports)
}
