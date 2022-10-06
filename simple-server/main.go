package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	log.SetOutput(os.Stdout)

	// Register handler function.
	mux := http.NewServeMux()
	mux.HandleFunc("/healthz", ping)

	// Wrap serverMux into a loggedMux and start serving.
	loggedMux := &logMux{mux}
	log.Fatal(http.ListenAndServe("localhost:80", loggedMux))
}

// ping is a handler used to check server health info
func ping(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(http.StatusOK)
}
