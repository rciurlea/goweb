package main

import (
	"log"
	"net/http"
)

func main() {
	log.SetFlags(log.Lmicroseconds)
	log.Println("Starting server on port 3000...")
	err := http.ListenAndServe(":3000", http.FileServer(http.Dir("static")))
	check(err)
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
