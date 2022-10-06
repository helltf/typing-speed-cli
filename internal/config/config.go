package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

var Conf *Config = readConfig()

type Config struct {
	Space string `json:"space"`
	Unit  string `json:"unit"`
}

func UpdateConfig(conf *Config) {
	Conf = conf
}

func readConfig() *Config {
	jsonFile, err := os.Open("./config.json")

	if err != nil {
		fmt.Println(err)
	}

	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)
	var result Config

	err = json.Unmarshal([]byte(byteValue), &result)

	if err != nil {
		panic(err)
	}

	return &result
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

	return writeConfig()
}

func SetUnit(unit string) error {
	Conf.Unit = unit

	return writeConfig()
}
