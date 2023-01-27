package main

import (
	anonymousfunction "learngo/anonymous_function"
	"learngo/array"
	"learngo/boolean"
	"learngo/closure"
	continuebreak "learngo/continue_break"
	"learngo/data_types"
	errors "learngo/errors"
	fr "learngo/for"
	"learngo/function"
	"learngo/goroutines"
	"learngo/helloworld"
	ifelse "learngo/if_else"
	ie "learngo/interface"
	"learngo/operators"
	pe "learngo/package"
	"learngo/pointer"
	pointerstruct "learngo/pointer_struct"
	printoutput "learngo/print_output"
	re "learngo/range"
	"learngo/recursion"
	"learngo/slice"
	"learngo/string"
	st "learngo/struct"
	sw "learngo/switch"

	deferpanicrecover "learngo/defer_panic_recover"
	pointerfunction "learngo/pointer_function"
	takeinput "learngo/take_input"
	typeassertions "learngo/type_assertions"
	typecasting "learngo/type_casting"
	variablescope "learngo/variable_scope"
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
	function.Function()
	variablescope.VariableScope()
	recursion.Recursion()
	anonymousfunction.AnonymousFunction()
	closure.Closure()
	pe.Package()
	pointer.Pointer()
	pointerfunction.PointerFunction()
	pointerstruct.PointerStruct()
	ie.Interface()
	ie.MultipleStruct()
	ie.EmptyInterface()
	typeassertions.TypeAssertions()
	errors.Error()
	// deferpanicrecover.Defer()
	// deferpanicrecover.Panic()
	deferpanicrecover.Recover()
	goroutines.Goroutines()
	goroutines.Ex()
}
