package main

import (
	"fmt"
	"os"
	"reflect"
	"strings"
)

var funcMap = map[string]interface{}{
	"GCS_CCC_OS_C1_TR01": GCS_CCC_OS_C1_TR01,
	"GCS_CCC_OS_C1_TR02": GCS_CCC_OS_C1_TR02,
	"GCS_CCC_OS_C1_TR03": GCS_CCC_OS_C1_TR03,
}

// Pass takes any number of strings and prints them as a single string
func CFIPass(values ...string) {
	fmtPrint(values, "PASS")
}

// FAIL takes any number of strings and prints them as a single string
func CFIFAIL(values ...string) {
	fmtPrint(values, "FAIL")
}

// Error takes any number of strings and prints them as a single string
func CFIError(values ...string) {
	fmtPrint(values, "ERROR")
}

// fmtPrint prints the provided values as a single string with a prefix
func fmtPrint(values []string, prefix string) {
	str := strings.Join(values, " ")
	fmt.Printf("[%s]: %s\n", prefix, str)
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
