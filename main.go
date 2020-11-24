package main

import (
    "fmt"
    "net/http"
    "log"
)

var Version = "dev"

func main() {
    http.HandleFunc("/", HelloServer)

    log.Println("Listening @:8080")
    http.ListenAndServe(":8080", nil)
}

func HelloServer(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, Version)
}
