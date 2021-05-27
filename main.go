package main

import (
	"fmt"
	"net/http"
	"runtime"
)

func osInfoEndpoint(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "server is running on %s\n", runtime.GOOS)
}

func main() {
	http.HandleFunc("/", osInfoEndpoint)
	fmt.Println("running server...")
	http.ListenAndServe(":8080", nil)
}
