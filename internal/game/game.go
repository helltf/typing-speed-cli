package game

import (
	"strings"

	"github.com/TwiN/go-color"
	"github.com/helltf/typing-speed-cli/internal/config"
	"github.com/helltf/typing-speed-cli/internal/writer"
)

type Game struct {
	context      string
	currentIndex int
	contextSlice []rune
	writer       *writer.Writer
}

func NewGame(context string) *Game {
	writer := writer.NewWriter()
	game :=&Game{
		context:      context,
		currentIndex: 0,
		contextSlice: []rune(context),
		writer:       writer} 
		
	writer.Print(game.getOutputContext())

	return game
}

func (game *Game) Input(input rune) bool {
	isCorrect := game.IsCorrectLetter(input)

	if !isCorrect {
		return false
	}

	game.setIndex(game.currentIndex + 1)

	game.writer.Update(game.getOutputContext())

	return game.currentIndex == len(game.context)
}

func (game Game) IsCorrectLetter(letter rune) bool {
	currentLetter := game.contextSlice[game.currentIndex]
	if letter == currentLetter {
		return true
	}

	return letter == 0 && currentLetter == 32
}

func (game *Game) setIndex(index int) {
	game.currentIndex = index
}

func (g *Game) colorizeContext() string {

	begin := color.Green + string(g.contextSlice[:g.currentIndex]) + color.Reset
	end := string(g.contextSlice[g.currentIndex:len(g.contextSlice)])

	return begin + end
}

func (g *Game) Stop() {
	g.writer.Stop()
}

func(g *Game) getOutputContext() string{
	return strings.ReplaceAll(g.colorizeContext()," ",config.Conf.Space)
}