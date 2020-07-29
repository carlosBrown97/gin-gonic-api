package main

import (
	"gin-gonic-api/src"
	"log"
	"net/http"
)

func main() {
	r := src.Routes()

	log.Fatal(http.ListenAndServe(":3000", r))
}
