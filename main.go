package main

import (
	"fmt"
	"net/http"
)

func main() {
	port := "8090"
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "This is version 1\n")
	})
	http.ListenAndServe(":"+port, nil)
}
