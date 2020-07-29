package parser

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

func WriteExample(fileName string) {
	writeJSON := jHolderWriter{"Apple", int64(100), "world", false}

	readyJSON, err := json.Marshal(writeJSON)
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println("\n========================Writing Example==========================")
	err = ioutil.WriteFile(fileName, readyJSON, 0755)
	fmt.Println("Writing JSON: ", string(readyJSON))
	fmt.Println("\n=================================================================")

	if err != nil {
		fmt.Println(err.Error())
	}
}
