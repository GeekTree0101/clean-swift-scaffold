package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "clean-swift-scaffold",
	Short: "clean-swift-scaffold is a Clean-Swift source & test code auto generator",
	Long:  fmt.Sprintf("\033[32m%s\033[0m\n", logo),
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
