package game

import (
	"encoding/json"
	"io/ioutil"

	"github.com/helltf/typing-speed-cli/internal/util"
)

type Stats struct {
	Last    StatsData `json:"last"`
	Average StatsData `json:"average"`
}

type StatsData struct {
	Time  int     `json:"time"`
	Words int     `json:"words"`
	Cps   float64 `json:"cps"`
}

func GenerateStats(game *Game) *Stats {
	return &Stats{
		Last: StatsData{
			Time:  game.time,
			Words: game.words,
			Cps:   game.Cps,
		},
		Average: getAverageStats(game),
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

func getAverageStats(game *Game) StatsData {
	return StatsData{
		Time:  1,
		Words: 1,
		Cps:   1,
	}
}
