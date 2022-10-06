package game

import (
	"strconv"
	"strings"
	"time"

	"github.com/TwiN/go-color"
	"github.com/helltf/typing-speed-cli/internal/config"
	"github.com/helltf/typing-speed-cli/internal/writer"
)

const updateCycle = 250

var ticker = time.NewTicker(updateCycle * time.Millisecond)
var quit = make(chan struct{})

type Game struct {
	context      string
	currentIndex int
	contextSlice []rune
	time         int
	Cps          float64
	words        int
}

func NewGame(context string) *Game {
	game := &Game{
		context:      context,
		currentIndex: 0,
		contextSlice: []rune(context),
		words:        1}

	writer.Print(game.getOutputContext())
	go game.startTimer()
	return game
}

func (game *Game) Input(input rune) bool {
	if !game.IsCorrectLetter(input) {
		return false
	}

	game.incrementIndex()

	if input == 0 {
		game.updateWordCount()
	}

	writer.Update(game.getOutputContext())

	return game.currentIndex == len(game.context)
}

func (g *Game) incrementIndex() {
	g.setIndex(g.currentIndex + 1)

}

func (g *Game) startTimer() {
	go func() {
		for {
			select {
			case <-ticker.C:
				g.updateTime()
			case <-quit:
				ticker.Stop()
				return
			}
		}
	}()
}

func (g *Game) updateTime() {
	g.time += updateCycle

	g.Cps = float64(g.currentIndex) / (float64(g.time) / float64(1000))
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
	close(quit)
	writer.Stop()
}

func (g *Game) updateWordCount() {
	g.words += 1
}

func (g *Game) getOutputContext() string {
	return strings.ReplaceAll(g.colorizeContext(), " ", config.Conf.Space) +
		"\n\n" +
		strconv.Itoa(int(g.Cps)) +
		" Characters per second" +
		"\n" +
		strconv.Itoa(g.words) +
		"/" +
		strconv.Itoa(len(strings.Split(g.context, " "))) +
		" words"
}
