package main

import (
	"fmt"
	"log"
	"net/http"
)

var version = "dev"

func main() {
	http.HandleFunc("/", HelloServer)

	log.Printf("Version `%s` Listening @:8080", version)
	http.ListenAndServe(":8080", nil)
}

// HelloServer main handler
func HelloServer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, version)
}
