package main

import (
	"compliant-financial-infrastructure/cfi"
	"fmt"
	"os"
	"reflect"
)

func FunctionA() {
	cfi.Pass("This is", "a test")
}

func FunctionB() {
}

func FunctionC() {
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("[ERROR] Please provide a function name to call")
		return
	}

	funcName := os.Args[1]

	funcMap := map[string]interface{}{
		"FunctionA": FunctionA,
		"FunctionB": FunctionB,
		"FunctionC": FunctionC,
	}

	if function, exists := funcMap[funcName]; exists {
		reflect.ValueOf(function).Call(nil)
	} else {
		fmt.Print("[ERROR] Function does not exist")
		// panic and print function name if it does not exist
		panic(fmt.Sprintf("[ERROR] Function %s does not exist", funcName))
	}
}
