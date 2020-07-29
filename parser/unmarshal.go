package parser

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"reflect"
)

func ReadExample(fileName string) {
	file, err := ioutil.ReadFile(fileName)

	if err != nil {
		fmt.Println("Error reading", err.Error())
		return
	}
	fmt.Println("====================Reader Example======================")
	fmt.Println("JSON before parsing")
	fmt.Println("Type: ", reflect.TypeOf(file))
	fmt.Println("Int Value: ", file)
	fmt.Println()
	fmt.Println("Stringified Value: ", string(file))
	fmt.Println()

	readJSON := jHolder{}

	err = json.Unmarshal(file, &readJSON)

	if err != nil {
		fmt.Println("Malformed JSON", err.Error())
		return
	}
	fmt.Println("JSON after parsing")
	fmt.Println("Type", reflect.TypeOf(readJSON))
	fmt.Printf("%+v", readJSON)
	fmt.Println("")
	fmt.Println("========================================================")

}
