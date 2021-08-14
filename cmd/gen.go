package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	genCmd = &cobra.Command{
		Use:   "gen",
		Short: "gnerate spy",
		Run: func(cmd *cobra.Command, args []string) {

			if len(name) == 0 {
				fmt.Println("invalid usecase name")
				return
			}

			if len(sourcePath) == 0 {
				fmt.Println("invalid source filepath")
				return
			}

			if len(testPath) == 0 {
				fmt.Println("invalid test filepath")
				return
			}

			// usecases := strings.Split(usecasesString, ",")

		},
	}
	name           string
	usecasesString string
	sourcePath     string
	testPath       string
)

func init() {
	genCmd.Flags().StringVarP(&name, "name", "n", "", "Usecase name, ex: -n Detail or --name Detail")
	genCmd.Flags().StringVarP(&usecasesString, "usecases", "u", "", "Behavior names, ex: -u Reload,Next or --usecases Reload,Next")
	genCmd.Flags().StringVarP(&sourcePath, "source", "s", "", "Source Dir, ex: -s ./Projects or --source ./Projects")
	genCmd.Flags().StringVarP(&testPath, "test", "t", "", "Test Dir, ex: -t ./Projects or --test ./Projects")
	rootCmd.AddCommand(genCmd)
}
