package main

import (
	"net/http"

	"github.com/goji/httpauth"
)

func main() {
	finalHandler := http.HandlerFunc(final)
	authHandler := httpauth.SimpleBasicAuth("user", "pass")

	http.Handle("/", authHandler(finalHandler))
	http.ListenAndServe(":3000", nil)
}

func final(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("HEH!"))
}
