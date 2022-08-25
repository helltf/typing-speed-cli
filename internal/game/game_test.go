package game

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewGame(t *testing.T) {
	context := "exampleContext"
	contextSlice := strings.Split(context, "")
	game := NewGame(context)
	defaulIndex := 0

	assert := assert.New(t)

	assert.Equal(game.context, context, "Expected game context to be same as given context")
	assert.Equal(game.currentIndex, defaulIndex, "Expected index to be 0 at the start")
	assert.Equal(game.contextSlice, contextSlice, "Expected slice to be the same as context")
}

func TestInputReturnNoEnd(t *testing.T) {
	assert := assert.New(t)
	context := "exampleContext"
	game := NewGame(context)

	end := game.Input("a")

	assert.Equal(end, false, "End should not have ended on input")
}

func TestInputIndexIsLastPositionReturEnd(t *testing.T) {
	assert := assert.New(t)
	context := "a"
	game := NewGame(context)

	end := game.Input("a")

	assert.Equal(end, true, "End should have ended on that input")
}

func TestIsCorrectLetterDoesNotMatchReturnFalse(t *testing.T) {
	assert := assert.New(t)
	context := "context"
	game := NewGame(context)
	input := "f"

	isCorrect := game.IsCorrectLetter(input)
	assert.False(isCorrect, "Letter should not be correct")
}

func TestIsCorrectLetterDoesMatchReturnTrue(t *testing.T) {
	assert := assert.New(t)
	context := "context"
	game := NewGame(context)
	input := "c"

	isCorrect := game.IsCorrectLetter(input)
	assert.Truef(isCorrect, "Letter should be %v but received %v", "c", input)
}

func TestIsCorrectLetterDifferentIndeDoesNotMatchReturnFalse(t *testing.T) {
	assert := assert.New(t)
	context := "context"
	game := NewGame(context)
	game.currentIndex = 1
	input := "f"

	isCorrect := game.IsCorrectLetter(input)
	assert.False(isCorrect, "Letter should not be correct")
}
