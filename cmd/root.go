package cmd

import (
	"fmt"
	"os"

	"github.com/parsaaes/netwolf/cmd/server"
	"github.com/parsaaes/netwolf/config"
	"github.com/spf13/cobra"
)

// ExitFailure status code
const ExitFailure = 1

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	cfg := config.Init("config.yaml")

	var root = &cobra.Command{
		Use:   "nw",
		Short: "netwolf",
	}

	server.Register(root, cfg)

	if err := root.Execute(); err != nil {
		fmt.Printf("error: %s\n", err.Error())
		os.Exit(ExitFailure)
	}
}
