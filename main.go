package main

import (
 "fmt"
 "net/http"
)

func main() {
 http.HandleFunc("/", HelloServer)
 fmt.Println("The server is listening...")
 fmt.Println("on localhost:8080")
 http.ListenAndServe(":8080", nil)
}

func HelloServer(writer http.ResponseWriter, request *http.Request) {
 fmt.Fprintf(writer, "Hello!")
}
