package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/greet", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello World!")
	})

	// ListenAndServe always returns a non-nil error
	// Fatal is equivalent to Print() followed by a call to os.Exit(1)
	//https://pkg.go.dev/log#Fatal
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}
