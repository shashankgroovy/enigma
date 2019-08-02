package utils

import (
	"log"
	"net/http"
)

func RequestLogger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		log.Printf(`%s %s`, r.Method, r.RequestURI)
		next.ServeHTTP(w, r)
	})
}
