package main

import (
	"log"
	"net/http"
)

func main() {
	log.SetFlags(log.Lmicroseconds)
	log.Println("Starting server on port 3000...")
	err := http.ListenAndServe("localhost:3000", http.HandlerFunc(staticServer))
	check(err)
}

func staticServer(w http.ResponseWriter, r *http.Request) {
	log.Println("Got request for", r.URL)
	http.ServeFile(w, r, "static")
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
