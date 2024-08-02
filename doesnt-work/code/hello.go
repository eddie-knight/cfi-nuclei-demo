package main

import (
	"fmt"
	"os"
	"reflect"
)

// var ctx = context.Background()

// // function to connect to GCP
// func connectToGCP() {
// 	client, err := storage.NewClient(ctx)
// 	if err != nil {

// 	}
// 	return client
// }

// func getBucket() {
// 	// connect to GCP
// 	client = connectToGCP()
// 	bucket := client.Bucket("my-bucket")
// }

func FunctionA() {
	fmt.Println("FunctionA called")
}

func FunctionB() {
	fmt.Println("FunctionB called")
}

func FunctionC() {
	fmt.Println("FunctionC called")
}

func fail() {
	fmt.Println("Fail")
}

func pass() {
	fmt.Println("Pass")
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
