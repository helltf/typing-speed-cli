/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"strings"

	"github.com/helltf/typing-speed-cli/internal/config"
	"github.com/spf13/cobra"
)

// setCmd represents the set command
var setCmd = &cobra.Command{
	Use:   "set",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		space, _ := cmd.Flags().GetString("space")

		if strings.ToLower(space) == "default"{
			space = " a"
		}

		err := setSpaceChar(space)

		if err != nil {
			panic(err)
		}

		fmt.Println("Successfully set space character")
	},
}

func init() {
	configCmd.AddCommand(setCmd)
	setCmd.PersistentFlags().String("space", "", "Set your space character")
}

func setSpaceChar(char string) error {
	err := config.SetSpace(char)

	return err
}