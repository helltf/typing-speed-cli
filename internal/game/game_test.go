package game

import "testing"

func TestNewGame(t *testing.T) {
	context := "exampleContext"
	game := NewGame(context)

	if game.context != context {
		t.Error("Expected game context to be same as given context")
	}
}
