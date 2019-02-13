package main

import (
	"debate-api/opinion"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc(opinion.ENDPOINT, opinion.Handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
