package main

import (
	"fmt"
	"net/http"
)

func SampleHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, net/http!")
}
