package main

import (
	"bytes"
	"net/http"
)

func enforceXMLHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// check for request body
		if r.ContentLength == 0 {
			http.Error(w, http.StatusText(400), 400)
			return
		}
		// sniff MIME type
		buf := new(bytes.Buffer)
		buf.ReadFrom(r.Body)
		if http.DetectContentType(buf.Bytes()) != "text/xml; charset=utf-8" {
			http.Error(w, http.StatusText(415), 415)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func main() {
	http.Handle("/", enforceXMLHandler(http.HandlerFunc(final)))
	http.ListenAndServe(":3000", nil)
}

func final(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("YAY!\n"))
}
