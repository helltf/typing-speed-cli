package game

import (
	"strings"
)

type Game struct {
	context      string
	currentIndex int
	contextSlice []string
}

func NewGame(context string) *Game {
	return &Game{
		context:      context,
		currentIndex: 0,
		contextSlice: strings.Split(context, "")}
}

func (game Game) Input(input string) bool {
	if game.contextSlice[len(game.contextSlice)-1] == input {
		return true
	}

	return false
}

func (game Game) IsCorrectLetter(letter string) bool {
	if letter == game.contextSlice[game.currentIndex] {
		return true
	}

	return false
}
