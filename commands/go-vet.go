package commands

import (
	"github.com/spf13/cobra"
)

var goVet = &cobra.Command{
	Use:   "go-vet [files...]",
	Short: "Check your Go source with go vet",
	Long:  ``,
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		runTool("go", append([]string{"vet"}, pkgNames(dirNames(args))...))
	},
}

func init() {
	rootCmd.AddCommand(goVet)
}
