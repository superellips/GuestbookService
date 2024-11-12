package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/guestbook/create", CreateHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
