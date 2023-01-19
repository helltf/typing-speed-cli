package config

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"

	"github.com/helltf/typing-speed-cli/internal/enum/unit"
	"github.com/helltf/typing-speed-cli/internal/util"
)

var Conf *Config = readConfig()

type Config struct {
	Space    string `json:"space"`
	Unit     string `json:"unit"`
	Cursor   bool   `json:"cursor"`
	Language string `json:"lang"`
}

var valid_units = []string{unit.Cps, unit.Wpm, unit.Cpm}
var valid_languages = []string{"en", "de", "es", "fr"}

func UpdateConfig(conf *Config) {
	Conf = conf
}

func readConfig() *Config {
	path := "./config.json"

	result, err := util.ReadJsonFile[Config](path)

	if err != nil {
		log.Fatal(err)
	}

	return result
}

func writeConfig() error {
	file, err := json.MarshalIndent(Conf, "", " ")

	if err != nil {
		return err
	}

	return ioutil.WriteFile("config.json", file, 0644)
}

func SetSpace(char string) error {
	Conf.Space = char

	return nil
}

func SetUnit(unit string) error {
	if !util.Contains(valid_units, unit) {
		return errors.New("invalid Unit")
	}

	Conf.Unit = unit

	return nil
}

func SetCursor(cursor bool) error {
	Conf.Cursor = cursor

	return nil
}

func SetLanguage(language string) error {
	if !util.Contains(valid_languages, language) {
		return errors.New("Invalid language")
	}

	Conf.Language = language

	return nil
}

func Write() error {
	return writeConfig()
}
