package commands

import (
	"github.com/spf13/cobra"
)

var goModTidy = &cobra.Command{
	Use:   "go-mod-tidy",
	Short: "Prune no-longer required commits from go.mod",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		runTool("go", append([]string{"mod", "tidy"}))
	},
}

func init() {
	rootCmd.AddCommand(goModTidy)
}
