/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"strconv"

	"github.com/helltf/typing-speed-cli/internal/game"
	"github.com/spf13/cobra"
)

// statsCmd represents the stats command
var statsCmd = &cobra.Command{
	Use:   "stats",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		stats := game.ReadStats()

		fmt.Println("Your Stats \n" +
			"Words in last run: " +
			strconv.Itoa(stats.Last.Words) + "\n" +
			"time spend in last run: " + strconv.Itoa(stats.Last.Time) + "s\n" +
			"characters per second: " + strconv.Itoa(int(stats.Last.Cps)))
	},
}

func init() {
	rootCmd.AddCommand(statsCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// statsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// statsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
