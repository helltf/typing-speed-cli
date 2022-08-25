/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/eiannone/keyboard"
	"github.com/helltf/typing-speed-cli/internal/game"
	"github.com/spf13/cobra"
)

// playCmd represents the play command
var playCmd = &cobra.Command{
	Use:   "play",
	Short: "Play a game to test you typing speed",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		context := "This is an example context"
		runningGame := game.NewGame(context)

		keyStrokes := make(chan string)
		go getKeys(keyStrokes)

		for key := range keyStrokes {
			runningGame.Input(key)
		}
	},
}

func init() {
	rootCmd.AddCommand(playCmd)
}

func getKeys(c chan<- string) {
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
			close(c)
			break
		}

		c <- string(char)
	}
}
