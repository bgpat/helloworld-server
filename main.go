package main

import (
	"io"
	"log"
	"net/http"
)

type Handle struct{}

func (h *Handle) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "hello, world.")
	log.Printf("%s\t%s\n", r.Method, r.URL)
}

func main() {
	http.ListenAndServe(":8080", &Handle{})
}
