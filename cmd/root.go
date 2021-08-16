package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "clean-swift-scaffold",
	Short: "clean-swift-scaffold is a Clean-Swift source & test code auto generator",
	Long:  "clean-swift-scaffold is a Clean-Swift source & test code auto generator",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
