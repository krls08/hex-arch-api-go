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
	mux.HandleFunc("/health", healthHandler)

	log.Fatal(http.ListenAndServe(httpAddr, mux))
}

func healthHandler(w http.ResponseWriter, _ *http.Request) {
	fmt.Println("healthHanlder -> start")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("everything is ok!"))
}
