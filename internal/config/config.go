package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Config struct {
	Space string`json:"space"`
}

func readConfig() Config {
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

	return result
}

func New() *Config{
	config := readConfig()
	return &config
}

func (c *Config) getSpace() string{
	return c.Space
}