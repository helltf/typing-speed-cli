package context

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Context struct {
	id string `json:"id"`
	context string `json:para`
}

type ContextNodes struct {
	data []Context `json:data`
}

func GetContext() string {
	context := readContext()
	fmt.Println(context)
	return context.data[0].context
}

func readContext() ContextNodes {
	jsonFile, err := os.Open("./data/context.json")

	if err != nil {
		fmt.Println(err)
	} 

	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)
    var result ContextNodes

    err = json.Unmarshal([]byte(byteValue), &result)

	if err != nil {
		panic(err)
	}
	return result
}
