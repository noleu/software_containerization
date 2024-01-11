package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	// set the flags for the logging package to give the filename in the logs
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	log.Println("starting sever")
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = fmt.Fprintln(w, "Hello World")
	})
	log.Fatal(http.ListenAndServe(":8080", nil))
}
