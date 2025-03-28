package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/gorilla/schema"
)

type Bill struct {
	Serial int `form:"serial"`
	Price  int `form:"price"`
	Units  int `form:"units"`
	Total  int `form:"total"`
}

type Bills struct {
	AllBills []Bill `form:"Bills"`
}

func main() {

	// registering the endpoints with default server
	http.HandleFunc("/", form)

	// starting the server
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func form(w http.ResponseWriter, r *http.Request) {

	fmt.Println("Method used is: ", r.Method)

	if r.Method == "GET" {
		//if method is get, just send the form, along with the data if possible
		t, _ := template.ParseFiles("./html/form.htm")
		t.Execute(w, "login")
	}
	if r.Method == "POST" {
		//if method is post, just print the values in the terminal
		//I know that the data is going to be a bill type thing in a form
		//just print it
		var values = new(Bills)
		err := r.ParseForm()
		if err != nil {
			fmt.Println("parseform:", err)
		}
		fmt.Println("form", r.Form)

		decoder := schema.NewDecoder()
		err = decoder.Decode(values, r.Form)
		if err != nil {
			fmt.Println("decode", err)
		}

		fmt.Println("Values:", values)
	}

	//will send the login page
}
