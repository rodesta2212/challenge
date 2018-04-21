package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Halaman Beranda", r.URL.Path)
	})

	http.ListenAndServe("localhost:8282", nil)
}
