package main

import (
	"fmt"
	"log"
	"net/http"
)

const serverAddress = "localhost:8093"

func main() {
	fmt.Println("http server is available at", serverAddress)
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(serverAddress, nil))
}

func handler(w http.ResponseWriter, req *http.Request) {
	fmt.Println("OK", req.Method)
	w.Write([]byte("OK"))
}
