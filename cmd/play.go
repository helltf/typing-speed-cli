/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/eiannone/keyboard"
	"github.com/helltf/typing-speed-cli/internal/context"
	"github.com/helltf/typing-speed-cli/internal/game"
	"github.com/spf13/cobra"
)

var playCmd = &cobra.Command{
	Use:   "play",
	Short: "Play a game to test you typing speed",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		start()
	},
}

func init() {
	rootCmd.AddCommand(playCmd)
}

func start() {
	gameContext := context.GetContext()
	runningGame := game.NewGame(gameContext)

	startKeyboard(runningGame)
}

func startKeyboard(game *game.Game) {
	if err := keyboard.Open(); err != nil {
		panic(err)
	}
	defer func() {
		_ = keyboard.Close()
	}()

	for {
		char, key, err := keyboard.GetKey()
		if err != nil {
			panic(err)
		}

		if key == keyboard.KeyEsc {
			break
		}

		isFinished := game.Input(char)

		if isFinished {
			game.Stop()
			break
		}
	}
}
