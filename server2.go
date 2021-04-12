package main

import (
	"fmt"
	"net/http"
)

type Person struct {
	Name string `json:"name"`
}

func Main2() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "%sさんの運勢は%sです", r.FormValue("p"), MakeFortune())
	})

	http.ListenAndServe(":8080", nil)
}
