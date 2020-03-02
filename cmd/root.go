package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// ExitFailure status code
const ExitFailure = 1

func Execute() {
	var root = &cobra.Command{
		Use:   "nw",
		Short: "netwolf",
		Run: func(cmd *cobra.Command, args []string) {

		},
	}

	if err := root.Execute(); err != nil {
		fmt.Printf("error: %s\n", err.Error())
		os.Exit(ExitFailure)
	}
}
