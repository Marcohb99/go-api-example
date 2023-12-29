package main

import (
	"fmt"
	"log"
	"net/http"
)

const httpAddr = ":8080"

func main() {
	fmt.Println("Server running on", httpAddr)

	mux := http.NewServeMux()

	// ENDPOINTS

	// hc
	mux.HandleFunc("/health", healthHandler)

	// ListenAndServe returns error
	log.Fatal(http.ListenAndServe(httpAddr, mux))
}

// 2nd param ignore because hc endpoints do not need request data
func healthHandler(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("everything is ok!"))
}
