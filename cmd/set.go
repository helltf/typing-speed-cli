/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"errors"
	"fmt"
	"strings"

	"github.com/helltf/typing-speed-cli/internal/config"
	"github.com/helltf/typing-speed-cli/internal/util"
	"github.com/spf13/cobra"
)

const (
	spaceConfig = "space"
	unitConfig  = "unit"
)

var valid_units = []string{"cps", "wpm"}

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
		space, _ := cmd.Flags().GetString(spaceConfig)
		unit, _ := cmd.Flags().GetString(unitConfig)

		if space != "" {
			err :=	updateSpaceChar(space)

			if err != nil {
				panic(err)
			}
		}

		if unit != "" {
			err := setSpeedUnit(unit)

			if err != nil {
				panic(err)
			}
		}

		fmt.Println("Successfully updated config")
	},
}

func updateSpaceChar(space string) error {
	if strings.ToLower(space) == "default" {
		space = " a"
	}

	err := setSpaceChar(space)
	return err
}

func init() {
	configCmd.AddCommand(setCmd)
	setCmd.PersistentFlags().String(spaceConfig, "", "Set your space character")
	setCmd.PersistentFlags().String(unitConfig, "", "Set your desired typing unit")
}

func setSpeedUnit(unit string) error {
	if !util.Contains(valid_units, unit) {
		return errors.New("invalid Unit")
	}

	err := config.SetUnit(unit)

	if err != nil {
		return err
	}

	return nil
}

func setSpaceChar(char string) error {
	err := config.SetSpace(char)

	return err
}
