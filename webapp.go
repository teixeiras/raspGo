package main

import (
    "fmt"
    "net/http"    
    "time"
)

func log(message string) {
	t := time.Now()
    fmt.Printf("%s: %s\n", t, message)
	
}

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

func main() {

	log("Server has started")

    http.HandleFunc("/", handler)
    http.ListenAndServe(":9999", nil)
}