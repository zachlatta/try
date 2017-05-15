package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/zachlatta/try/runner"
)

var RootCmd = &cobra.Command{
	Use:   "try [url]",
	Short: "Try running a given project",
	Run: func(md *cobra.Command, args []string) {
		if len(args) == 0 || len(args) > 1 {
			fmt.Fprintln(os.Stderr, "Please pass one argument.")
			os.Exit(1)
		}

		repoUrl := args[0]

		if err := runner.Run(repoUrl); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	},
}

func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
