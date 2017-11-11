package main

import (
	"net/http"
	"log"
	"net/http/httputil"
	"net/url"
)

const backendServerURI = "https://postman-echo.com/"
const port = ":80"

func main() {
	log.Printf("CORS server on %s", port)
	uri, err := url.Parse(backendServerURI)
	if err != nil {
		log.Fatal("Invalid URI")
	}
	reverseProxy := httputil.NewSingleHostReverseProxy(uri)
	h := middlewareCORS(reverseProxy)
	h = middlewareSameHost(h)
	log.Fatal(http.ListenAndServe(port, h))
}

func middlewareCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "X-Requested-With")
		next.ServeHTTP(w, r)
	})
}
func middlewareSameHost(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		r.Host = r.URL.Host
		next.ServeHTTP(w, r)
	})
}
