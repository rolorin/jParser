package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"reflect"
)

type jHolder struct {
	Hello   string
	Int     int64
	String  string
	Boolean bool
}

type jHolderWriter struct {
	Hello   string `json:"hello"`
	Int     int64  `json:"integer"`
	String  string `json:"str"`
	Boolean bool   `json:"bool,omitempty"`
}

func ReadExample(fileName string) {
	file, err := ioutil.ReadFile(fileName)

	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("JSON before parsing")
	fmt.Println("Type", reflect.TypeOf(file))
	fmt.Println("Value", file)
	fmt.Println()

	readJSON := jHolder{}

	err = json.Unmarshal(file, &readJSON)

	fmt.Println("JSON after parsing")
	fmt.Println("Type", reflect.TypeOf(readJSON))
	fmt.Printf("%+v", readJSON)
	fmt.Println("")

	if err != nil {
		fmt.Println(err.Error())
		return
	}
}

func WriteExample(fileName string) {
	writeJSON := jHolderWriter{"Apple", int64(100), "world", false}

	readyJSON, err := json.Marshal(writeJSON)
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println("\nWriting Example")
	err = ioutil.WriteFile(fileName, readyJSON, 0755)
	fmt.Println("Writing JSON", string(readyJSON))

	if err != nil {
		fmt.Println(err.Error())
	}
}

func main() {

	if len(os.Args) < 2 {
		fmt.Println("Usage: ", "./main [file.json]")
		return
	}

	fileName := os.Args[1]
	ReadExample(fileName)

	if len(os.Args) > 2 {
		WriteExample(fileName)
	}
	ioutil.WriteFile("out.json", []byte("hello"), 0644)
	// file, _ := io.ReadFile("")
}
