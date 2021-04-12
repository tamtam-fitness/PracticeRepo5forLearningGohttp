package main

import (
	"fmt"
	"math/rand"
	"net/http"
)

func MakeFortune() string {

	num := rand.Intn(4)

	result := ""
	switch num {
	case 0:
		result = "大吉"
	case 1:
		result = "吉"

	case 2:
		result = "凶"

	default:
		result = "大凶"
	}
	return result

}

func fortunehandler(w http.ResponseWriter, r *http.Request) {
	result := MakeFortune()
	fmt.Fprint(w, result)
}

func Main1() {

	http.HandleFunc("/", fortunehandler)
	http.ListenAndServe(":8080", nil)

}
