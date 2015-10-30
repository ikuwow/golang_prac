package main

import (
    "fmt"
    "net/http" // web server
)

func handler( w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hi %s!", r.URL.Path[1:])
}

func main() {

    // if accessed to /, go up handler
    http.HandleFunc("/", handler)
    http.ListenAndServe(":8888", nil)
}
