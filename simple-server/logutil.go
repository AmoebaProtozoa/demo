package main

import (
	"log"
	"net/http"
	"os"
)

// logMux wraps regular server mux to provide logging functions.
type logMux struct {
	*http.ServeMux
}

// logWriter wraps http.ResponseWriter to record the status code.
type logWriter struct {
	http.ResponseWriter
	statusCode int
}

func (lw *logWriter) WriteHeader(statusCode int) {
	lw.statusCode = statusCode
	lw.ResponseWriter.WriteHeader(statusCode)
}

func (lm *logMux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// Copy whatever is in request header to response header.
	for k, vs := range r.Header {
		for _, v := range vs {
			w.Header().Add(k, v)
		}
	}
	versionInfo := os.Getenv("VERSION")
	w.Header().Set("version", versionInfo)

	lw := &logWriter{
		ResponseWriter: w,
		// Defaulting status code to 200 here since implicit WriteHeader
		// via Write will return a status code of 200.
		statusCode: http.StatusOK,
	}

	lm.ServeMux.ServeHTTP(lw, r)

	log.Printf("response status code: %d\t version: %s\t client address: %s\n",
		lw.statusCode,
		versionInfo,
		r.RemoteAddr,
	)
}
