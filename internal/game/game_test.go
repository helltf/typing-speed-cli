package game

import (
	"testing"

	"github.com/helltf/typing-speed-cli/internal/config"
	"github.com/stretchr/testify/assert"
)

func loadConfig(){
	mockedConfig := &config.Config{
		Space:       "-"}
	config.InitWithConf(mockedConfig)
}

func TestNewGame(t *testing.T) {
	loadConfig()
	context := "exampleContext"
	contextSlice := []rune(context)
	game := NewGame(context)
	defaulIndex := 0

	assert := assert.New(t)

	assert.Equal(game.context, context, "Expected game context to be same as given context")
	assert.Equal(game.currentIndex, defaulIndex, "Expected index to be 0 at the start")
	assert.Equal(game.contextSlice, contextSlice, "Expected slice to be the same as context")
}

func TestInputReturnNoEnd(t *testing.T) {
	loadConfig()
	assert := assert.New(t)
	context := "exampleContext"
	game := NewGame(context)

	end := game.Input([]rune("a")[0])

	assert.Equal(end, false, "End should not have ended on input")
}

func TestInputIndexIsLastPositionReturEnd(t *testing.T) {
	loadConfig()
	assert := assert.New(t)
	context := "a"
	game := NewGame(context)

	end := game.Input([]rune("a")[0])

	assert.Equal(end, true, "End should have ended on that input")
}

func TestIsCorrectLetterDoesNotMatchReturnFalse(t *testing.T) {
	loadConfig()
	assert := assert.New(t)
	context := "context"
	game := NewGame(context)
	input := []rune("f")

	isCorrect := game.IsCorrectLetter(input[0])
	assert.False(isCorrect, "Letter should not be correct")
}

func TestIsCorrectLetterDoesMatchReturnTrue(t *testing.T) {
	loadConfig()
	assert := assert.New(t)
	context := "context"
	game := NewGame(context)
	input := []rune("c")

	isCorrect := game.IsCorrectLetter(input[0])
	assert.Truef(isCorrect, "Letter should be %v but received %v", "102", input)
}

func TestIsCorrectLetterDifferentIndexDoesNotMatchReturnFalse(t *testing.T) {
	loadConfig()
	assert := assert.New(t)
	context := "context"
	game := NewGame(context)
	game.currentIndex = 1
	input := []rune("f")

	isCorrect := game.IsCorrectLetter(input[0])
	assert.False(isCorrect, "Letter should not be correct")
}
