package main

import (
	"net/http"
)

func main() {
	router := http.NewServeMux()
	router.HandleFunc("GET /", myHandler)
	http.ListenAndServe(":8080", router)
}
