package game

import (
	"fmt"
)

type Game struct {
	context      string
	currentIndex int
	contextSlice []rune
}

func NewGame(context string) *Game {
	fmt.Println(context)

	return &Game{
		context:      context,
		currentIndex: 0,
		contextSlice: []rune(context)}
}

func (game *Game) Input(input rune) bool {
	isCorrect := game.IsCorrectLetter(input)

	if !isCorrect {
		return false
	}

	game.setIndex(game.currentIndex + 1)

	fmt.Printf("%v\n", game.context[:game.currentIndex])

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
