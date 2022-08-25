package game

type Game struct {
	context string
}

func NewGame(context string) *Game {
	game := Game{context: context}

	return &game
}
