package main

import (
	"fmt"
	"os"
)

func main() {

	fmt.Println("Program name:", os.Args[0])
	if len(os.Args) > 1 {
		fmt.Println("Arguments Passed are: ")
		for i, arg := range os.Args {
			fmt.Printf("%d: %s\n", i, arg)
		}
	} else {
		fmt.Println("No arguments provided")
	}
}
