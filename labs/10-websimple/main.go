package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, James Bond. Anda tadi request %s\n", r.URL.Path)
	})

	http.ListenAndServe(":80", nil)
}
