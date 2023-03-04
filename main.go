package main

import (
	"fmt"
	"net/http"
)

func handlerfunc(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprintf(w, "<h1>Welcome to my awesome site!</h1>")
}

func main() {
	http.HandleFunc("/", handlerfunc)
	fmt.Println("Starting the server on :3000")
	http.ListenAndServe("localhost:3000", nil)
}