package game

import (
	"encoding/json"
	"io/ioutil"

	"github.com/helltf/typing-speed-cli/internal/util"
)

type Stats struct {
	Last    StatsData     `json:"last"`
	Average *StatsData    `json:"average"`
	History *StatsHistory `json:"history"`
}

type StatsHistory struct {
	Time  []int `json:"time"`
	Words []int `json:"words"`
	Cps   []int `json:"cps"`
}

type StatsData struct {
	Time  int `json:"time"`
	Words int `json:"words"`
	Cps   int `json:"cps"`
}

func GenerateStats(game *Game) *Stats {
	avr, history := getAverageStats(game)
	return &Stats{
		Last: StatsData{
			Time:  game.time,
			Words: game.words,
			Cps:   game.Cps,
		},
		Average: avr,
		History: history,
	}
}

func SaveStats(stats *Stats) error {
	file, err := json.MarshalIndent(stats, "", " ")

	if err != nil {
		return err
	}

	return ioutil.WriteFile("stats.json", file, 0644)
}

func ReadStats() (*Stats, error) {
	path := "stats.json"

	result, err := util.ReadJsonFile[Stats](path)

	if err != nil {
		return nil, err
	}

	return result, nil
}

func getAverageStats(game *Game) (*StatsData, *StatsHistory) {
	stats, err := ReadStats()

	if err != nil {
		return &StatsData{
				Time:  game.time,
				Words: game.words,
				Cps:   game.Cps,
			},
			&StatsHistory{
				Cps:   []int{},
				Time:  []int{},
				Words: []int{},
			}
	}
	timeAvr, timeHistory := getAverageAndHistory(stats.History.Time, game.time)
	cpsAvr, cpsHistory := getAverageAndHistory(stats.History.Cps, game.Cps)
	wordsAvr, wordsHistory := getAverageAndHistory(stats.History.Words, game.words)

	return &StatsData{
			Time:  timeAvr,
			Cps:   cpsAvr,
			Words: wordsAvr,
		},
		&StatsHistory{
			Words: wordsHistory,
			Cps:   cpsHistory,
			Time:  timeHistory,
		}
}

func getAverageAndHistory(slice []int, additional int) (int, []int) {
	history := append(slice, additional)
	sum := 0

	if len(history) >= 10 {
		history = history[1:]
	}

	for _, v := range history {
		sum += v
	}

	return sum / len(history), history
}
