package game

import (
	"strconv"
	"strings"
	"time"

	"github.com/TwiN/go-color"
	"github.com/helltf/typing-speed-cli/internal/config"
	"github.com/helltf/typing-speed-cli/internal/enum/unit"
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
	Cps          int
	words        int
}

func NewGame(context string) *Game {
	game := &Game{
		context:      context,
		currentIndex: 0,
		contextSlice: []rune(context),
		words:        1}

	writer.Print(game.getOutputContext(true))
	go game.startTimer()
	return game
}

func (game *Game) Input(input rune) bool {
	correctLetter := game.IsCorrectLetter(input)

	if correctLetter {
		game.incrementIndex()

		if input == 0 {
			game.updateWordCount()
		}
	}

	writer.Update(game.getOutputContext(correctLetter))

	if !correctLetter {
		return false
	}

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

	g.Cps = int(float64(g.currentIndex) / (float64(g.time) / float64(1000)))
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

func (g *Game) colorizeContext(correct bool) string {
	cursor := ""
	if config.Conf.Cursor {
		cursor = "|"
	}
	if !correct {
		cursor = color.Red + cursor + color.Reset
	}

	begin := color.Green + string(g.contextSlice[:g.currentIndex]) + color.Reset
	end := string(g.contextSlice[g.currentIndex:len(g.contextSlice)])

	return begin + cursor + end
}

func (g *Game) Stop() {
	close(quit)
	writer.Stop()
	SaveStats(GenerateStats(g))
}

func (g *Game) updateWordCount() {
	g.words += 1
}

func (g *Game) getCps() int {
	return int(g.Cps)
}

func (g *Game) getWpm() int {
	return g.getCpm() / 5
}

func (g *Game) getCpm() int {
	return g.getCps() * 60
}

func (g *Game) getCurrentWords() string {
	return strconv.Itoa(g.words)
}

func (g *Game) getMaxWords() string {
	return strconv.Itoa(len(strings.Split(g.context, " ")))
}

func (g *Game) getCurrentSpeed() string {
	if config.Conf.Unit == unit.Cps {
		return strconv.Itoa(g.getCps()) + " Characters per second"
	}

	if config.Conf.Unit == unit.Cpm {
		return strconv.Itoa(g.getCpm()) + " Characters per minute"
	}

	return strconv.Itoa(g.getWpm()) + " Words per minute"
}

func (g *Game) getOutputContext(correct bool) string {
	return strings.ReplaceAll(g.colorizeContext(correct), " ", config.Conf.Space) +
		"\n\n" +
		g.getCurrentSpeed() +
		"\n" +
		g.getCurrentWords() +
		"/" +
		g.getMaxWords() +
		" words"
}
