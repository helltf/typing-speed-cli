package util

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

func ReadJsonFile[T interface{}](path string) (*T, error) {
	jsonFile, err := os.Open(path)

	if err != nil {
		return nil, err
	}

	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)
	var result T

	err = json.Unmarshal([]byte(byteValue), &result)

	if err != nil {
		return nil, err
	}

	return &result, nil
}
