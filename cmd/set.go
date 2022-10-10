/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"errors"
	"fmt"
	"strings"

	"github.com/helltf/typing-speed-cli/internal/config"
	"github.com/helltf/typing-speed-cli/internal/enum/unit"
	"github.com/helltf/typing-speed-cli/internal/util"
	"github.com/spf13/cobra"
)

const (
	spaceConfig  = "space"
	unitConfig   = "unit"
	cursorConfig = "cursor"
)

var valid_units = []string{unit.Cps, unit.Wpm, unit.Cpm}

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
		cursor, _ := cmd.Flags().GetBool(cursorConfig)

		if space != "" {
			err := updateSpaceChar(space)

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

		err := setCursor(cursor)

		if err != nil {
			panic(err)
		}

		fmt.Println("Successfully updated config")
	},
}

func updateSpaceChar(space string) error {
	if strings.ToLower(space) == "default" {
		space = " "
	}

	err := setSpaceChar(space)
	return err
}

func init() {
	configCmd.AddCommand(setCmd)
	setCmd.PersistentFlags().String(spaceConfig, "", "Set your space character")
	setCmd.PersistentFlags().String(unitConfig, "", "Set your desired typing unit")
	setCmd.PersistentFlags().Bool(cursorConfig, config.Conf.Cursor, "enable/disable cursor")
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

func setCursor(cursor bool) error {
	return config.SetCursor(cursor)
}
