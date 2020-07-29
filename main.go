package main

import (
	"fmt"
	"os"

	"github.com/rolorin/jparser/parser"
)

func main() {

	if len(os.Args) < 2 {
		fmt.Println("Usage: ", "./main [input.json] ?[output.json]")
		return
	}

	fileName := os.Args[1]
	parser.ReadExample(fileName)
	parser.ArbitaryReadExample(fileName)
	if len(os.Args) > 2 {
		parser.WriteExample(os.Args[2])
	}
}
