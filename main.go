package main

import (
	"fmt"
	"go-examples/parallel"
	"go-examples/variadic"
	"os"
)

func main() {
	modules := map[string]func(){
		"variadic": variadic.Start,
		"parallel": parallel.Start}

	to_run := os.Args[1]

	fmt.Println("name to run: ", to_run)

	func_to_run := modules[to_run]

	func_to_run()
}
