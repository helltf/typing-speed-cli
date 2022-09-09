/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"github.com/helltf/typing-speed-cli/cmd"
	"github.com/helltf/typing-speed-cli/internal/config"
)

func main() {
	config.Init()
	cmd.Execute()
}
