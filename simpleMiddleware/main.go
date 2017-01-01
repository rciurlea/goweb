package main

import (
	"log"
	"net/http"
)

func middlewareOne(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("inside middleware 1...")
		next.ServeHTTP(w, r)
		log.Println("inside middleware 1 again...")
	})
}

func middlewareTwo(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("inside middleware 2...")
		if r.URL.Path != "/" {
			return
		}
		next.ServeHTTP(w, r)
		log.Println("inside middleware 2 again...")
	})
}

func final(w http.ResponseWriter, r *http.Request) {
	log.Println("inside final handler...")
	w.Write([]byte("OHAI"))
}

func main() {
	finalHandler := http.HandlerFunc(final)
	http.Handle("/", middlewareOne(middlewareTwo(finalHandler)))
	http.ListenAndServe(":3000", nil)
}
