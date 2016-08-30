package controllers

import (
	"net/http"
	"time"
	"log"
)

func HttpFilter(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		start := time.Now()
		log.Printf("Started %s %s ", req.Method, req.URL.Path)
		handler.ServeHTTP(w, req)
		log.Printf("Completed %s in %v", req.URL.Path, time.Since(start))
	})
}