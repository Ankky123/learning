package main

import (
	"fmt"
	"go-vanilla/config"
	"net/http"
)

const webPort = "8080"

func main() {

	app := config.Init()

	// let's register a function with the multiplexer
	http.HandleFunc("/", app.Greetings)
	http.HandleFunc("/getUsers", app.GetUsers)
	http.HandleFunc("/addUser", app.AddUser)
	// start a server
	fmt.Println("Starting a server on port", webPort)
	http.ListenAndServe(fmt.Sprintf(":%s", webPort), nil)

}
