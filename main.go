package main

import (
	"learngo/data_types"
	"learngo/helloworld"
	printoutput "learngo/print_output"
	takeinput "learngo/take_input"

	"learngo/variables"
)

func main() {
	helloworld.Hello()
	variables.Variables()
	data_types.DataTypes()
	takeinput.InputData()
	printoutput.PrintOutput()
}
