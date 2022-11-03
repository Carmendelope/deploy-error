package main

import (
    "fmt"
    "net/http"
)

func main() {
    http.HandleFunc("/", HelloServer)
    http.ListenAndServe(":8080", nil)
}

func HelloServer(w http.ResponseWriter, r *http.Request) {
    str := "This is a Golang example\nPlease, select a valid repository changing the LAUNCHER_TARGET_REPO value"
    https://github.com/Carmendelope/simple-app-go
    fmt.Fprintf(w, str)
}
