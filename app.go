package main

import (
	"net/http"
	"log"
)

func main() {
	mux := http.NewServeMux()
	log.Fatal(http.ListenAndServe(":2017", mux))
}
