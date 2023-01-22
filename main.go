package main

import (
	"learngo/array"
	"learngo/boolean"
	continuebreak "learngo/continue_break"
	"learngo/data_types"
	fr "learngo/for"
	"learngo/helloworld"
	ifelse "learngo/if_else"
	"learngo/operators"
	printoutput "learngo/print_output"
	re "learngo/range"
	"learngo/slice"
	"learngo/string"
	st "learngo/struct"
	sw "learngo/switch"
	takeinput "learngo/take_input"
	typecasting "learngo/type_casting"
	"learngo/while"

	mp "learngo/map"
	"learngo/variables"
)

func main() {
	helloworld.Hello()
	variables.Variables()
	data_types.DataTypes()
	takeinput.InputData()
	printoutput.PrintOutput()
	operators.ArithmeticOperator()
	operators.IncrementDecrementOperator()
	typecasting.TypeCasting()
	boolean.Boolean()
	ifelse.IfElse()
	sw.Switch()
	fr.For()
	while.While()
	re.Range()
	continuebreak.ContinueBreak()
	array.Array()
	array.MultidimensionalArray()
	slice.Slice()
	mp.Map()
	st.Struct()
	st.FunctionInsideStruct()
	string.String()
}
