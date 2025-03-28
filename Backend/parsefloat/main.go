package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	str := "15.5"

	val, err := strconv.ParseFloat(str, 64)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(str, val)

	envVar := os.Getenv("GOPATH")
	fmt.Println("GOPATH is ->", envVar)
	// fmt.Print(val == 1.2567)

	//let's create a file in the GOPATH src
	file, err := os.Create(envVar + "\\src\\Test.txt")
	if err != nil {
		log.Fatal("Some error occured :", err)
	}
	defer file.Close()

	file.WriteString("Some line in the file")

}
