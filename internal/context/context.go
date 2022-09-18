package context

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"time"
)

type Context struct {
	Id      json.Number `json:"id"`
	Context string      `json:"para"`
}

func GetContext() string {
	rand.Seed(time.Now().UnixNano())

	context := readContext()
	return context[rand.Intn(len(context))-1].Context
}

func readContext() []Context {
	jsonFile, err := os.Open("./data/context.json")

	if err != nil {
		fmt.Println(err)
	}

	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)
	var result []Context

	err = json.Unmarshal([]byte(byteValue), &result)

	if err != nil {
		panic(err)
	}

	return result
}
