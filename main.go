package main

import (
	"net/http"

	dd "x/usecase"
)

func main() {
	http.HandleFunc("/", dd.TemplateHandler)

	http.ListenAndServe(":3000", nil)
}




