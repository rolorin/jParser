package parser

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

func ArbitaryReadExample(fileName string) {
	file, err := ioutil.ReadFile(fileName)

	if err != nil {
		fmt.Println("Error reading", err.Error())
		return
	}
	fmt.Println("====================Arbitary Reader Example======================")

	var readJSON interface{}

	err = json.Unmarshal(file, &readJSON)

	if err != nil {
		fmt.Println("Malformed JSON", err.Error())
		return
	}
	arbitaryMap := readJSON.(map[string]interface{})

	for key, value := range arbitaryMap {
		switch actualValue := value.(type) {
		case string:
			fmt.Println(key, "is string:", actualValue)
		case float64:
			fmt.Println(key, "is float64:", actualValue)
		case bool:
			fmt.Println(key, "is bool:", actualValue)
		default:
			fmt.Println(key, "is of a type I don't know how to handle")
		}
	}

	fmt.Println("")
	fmt.Println("========================================================")

}
