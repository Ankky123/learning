package main

import (
	"log"
	"net/http"
	"time"
	"vanilla-go-serve/pkg/server"
)

func main() {

	address := "localhost:8080"
	mux := http.NewServeMux()

	srv := &server.Server{}

	mux.HandleFunc("/", srv.RespHome)
	mux.HandleFunc("/users", srv.RespUsers)

	s := &http.Server{
		Addr:           address,
		Handler:        mux,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	log.Printf("Started Server at: %v", address)
	log.Fatal(s.ListenAndServe())
}
