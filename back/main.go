package main

import (
	"log"
	"net/http"
)

func main() {
	var fs = http.FileServer(http.Dir("./front/html"))
	http.Handle("/", fs)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
