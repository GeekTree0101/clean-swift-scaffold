package cmd

import (
	"fmt"

	"github.com/Geektree0101/clean-swift-scaffold/internal/gen"
	"github.com/spf13/cobra"
)

var (
	genCmd = &cobra.Command{
		Use:   "run",
		Short: "generate source & unit tests files",
		Run: func(cmd *cobra.Command, args []string) {

			fmt.Printf("\033[32m%s\033[0m\n", logo)

			if len(name) == 0 {
				fmt.Println("[Error] invalid usecase name\033[0m")
				return
			}

			gen := gen.NewGenerator(
				gen.Genflag{
					Name:           name,
					UsecasesString: usecasesString,
					SourcePath:     sourcePath,
					TestPath:       testPath,
					ConfigFilePath: configFilePath,
				},
			)

			err := gen.Run()

			if err != nil {
				fmt.Printf("\033[31m[Error] failed to generate: %s\n\033[0m", err.Error())
			} else {
				fmt.Printf("\033[32m[Log] Done!\n\n\033[0m")
			}

		},
	}
	name           string
	usecasesString string
	sourcePath     string
	testPath       string
	configFilePath string
)

func init() {
	genCmd.Flags().StringVarP(&name, "name", "n", "", "Usecase name, ex: -n Detail or --name Detail")
	genCmd.Flags().StringVarP(&usecasesString, "usecases", "u", "", "Behavior names, ex: -u Reload,Next or --usecases Reload,Next")
	genCmd.Flags().StringVarP(&configFilePath, "config", "c", ".", "Configure file path, ex: -t ./Projects or --test ./Projects")
	// optional flags
	genCmd.Flags().StringVarP(&sourcePath, "source", "s", "", "Source dir, ex: -s ./Projects or --source ./Projects")
	genCmd.Flags().StringVarP(&testPath, "test", "t", "", "Test dir, ex: -t ./Projects or --test ./Projects")
	rootCmd.AddCommand(genCmd)
}
