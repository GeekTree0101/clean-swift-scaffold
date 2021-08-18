package clean_swift_scaffold

import (
	"fmt"

	"github.com/Geektree0101/clean-swift-scaffold/internal/gen"
	"github.com/spf13/cobra"
)

const (
	logo = `
	_____ _                  _____          _  __ _   
	/  __ \ |                /  ___|        (_)/ _| |  
	| /  \/ | ___  __ _ _ __ \ '--.__      ___| |_| |_ 
	| |   | |/ _ \/ _' | '_ \ '--. \ \ /\ / / |  _| __|
	| \__/\ |  __/ (_| | | | /\__/ /\ V  V /| | | | |_ 
	 \____/_|\___|\__,_|_| |_\____/  \_/\_/ |_|_|  \__|
																										 
																										 
	 _____            __  __      _     _              
	/  ___|          / _|/ _|    | |   | |             
	\ '--.  ___ __ _| |_| |_ ___ | | __| |             
	 '--. \/ __/ _' |  _|  _/ _ \| |/ _' |             
	/\__/ / (_| (_| | | | || (_) | | (_| |             
	\____/ \___\__,_|_| |_| \___/|_|\__._|

	Copyright © 2021 Geektree0101. All rights reserved.
	`
)

func NewRunnerCommand(use string) *cobra.Command {

	var name string
	var usecasesString string
	var sourceDir string
	var testDir string
	var configFilePath string

	genCmd := &cobra.Command{
		Use:   use,
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
					SourceDir:      sourceDir,
					TestDir:        testDir,
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

	genCmd.Flags().StringVarP(&name, "name", "n", "", "Usecase name, ex: -n Detail or --name Detail")
	genCmd.Flags().StringVarP(&usecasesString, "usecases", "u", "", "Behavior names, ex: -u Reload,Next or --usecases Reload,Next")
	genCmd.Flags().StringVarP(&configFilePath, "config", "c", "./config.yaml", "Configure file path, ex: -t ./some/config.yaml or --test ./some/config.yaml")
	// optional flags
	genCmd.Flags().StringVarP(&sourceDir, "source", "s", "", "Source dir, ex: -s ./Projects or --source ./Projects")
	genCmd.Flags().StringVarP(&testDir, "test", "t", "", "Test dir, ex: -t ./Projects or --test ./Projects")

	return genCmd
}
