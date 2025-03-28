package main

import (
	"fmt"
	"test-project/package1"
	"test-project/package2"
)

var SomeInt int

func main() {

	bottle := package1.Bottle{Name: "Milton", Material: "Plastic", Volume: 2}

	fmt.Println(bottle)
	package1.GetBottle(bottle)

	package2.IncreaseCapacity(&bottle)

	package1.GetBottle(bottle)
	//mux.NewRouter().ServeHTTP(http.ResponseWriter{}, *http.Request{})
	SomeInt = 1
	fmt.Println(SomeInt)
	fmt.Println("Creating the file and folder")

	package1.CreateFile()

}
