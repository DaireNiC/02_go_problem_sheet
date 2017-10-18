//Author : Daire Ni Chathain
// go exercise 1 : Guessing game

package main

import (
	"fmt"
	"net/http"
)

func requestHandler(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintln(w, "Guessing Game ")

}

func main() {
	http.HandleFunc("/", requestHandler)
	http.ListenAndServe(":8080", nil)
}
