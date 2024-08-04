package main

import (
	"fmt"
	"os"
	"reflect"
	"strings"
)

var funcMap = map[string]interface{}{
	"FunctionA": FunctionA,
	"FunctionB": FunctionB,
	"FunctionC": FunctionC,
}

func FunctionA() {
	CFIPass("This is a test")
}

func FunctionB() {
	CFIFail("This is a test")
}

func FunctionC() {
	CFIError("This is a test")
}

// func main() {
// 	if len(os.Args) < 2 {
// 		fmt.Println("[ERROR] Please provide a function name to call")
// 		return
// 	}

// 	funcName := os.Args[1]

// 	funcMap := map[string]interface{}{
// 		"FunctionA": FunctionA,
// 		"FunctionB": FunctionB,
// 		"FunctionC": FunctionC,
// 	}

// 	if function, exists := funcMap[funcName]; exists {
// 		reflect.ValueOf(function).Call(nil)
// 	} else {
// 		fmt.Print("[ERROR] Function does not exist")
// 		// panic and print function name if it does not exist
// 		panic(fmt.Sprintf("[ERROR] Function %s does not exist", funcName))
// 	}
// }

// Pass takes any number of strings and prints them as a single string
func CFIPass(values ...string) {
	fmtPrint(values, "PASS")
}

// FAIL takes any number of strings and prints them as a single string
func CFIFail(values ...string) {
	fmtPrint(values, "FAIL")
}

// Error takes any number of strings and prints them as a single string
func CFIError(values ...string) {
	fmtPrint(values, "ERROR")
}

// fmtPrint prints the provided values as a single string with a prefix
func fmtPrint(values []string, prefix string) {
	str := strings.Join(values, " ")
	fmt.Printf("[%s]: %s", prefix, str)
}

func main() {
	if len(os.Args) < 2 {
		CFIError("Please provide a function name to call")
		return
	}

	funcName := os.Args[1]
	if function, exists := funcMap[funcName]; exists {
		reflect.ValueOf(function).Call(nil)
	} else {
		CFIError("Function does not exist")
		// panic and print function name if it does not exist
		CFIError("Function %s does not exist", funcName)
	}
}
