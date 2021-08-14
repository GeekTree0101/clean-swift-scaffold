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

			if len(path) == 0 {
				fmt.Println("invalid filepath")
				return
			}

			// usecases := strings.Split(usecasesString, ",")

		},
	}
	name           string
	usecasesString string
	path           string
)

func init() {
	genCmd.Flags().StringVarP(&name, "name", "n", "", "-n Detail or --name Detail")
	genCmd.Flags().StringVarP(&usecasesString, "usecases", "u", "", "-u Reload,Next or --usecases Reload,Next")
	genCmd.Flags().StringVarP(&path, "path", "p", "", "-p ./Projects or --path ./Projects")
	rootCmd.AddCommand(genCmd)
}
