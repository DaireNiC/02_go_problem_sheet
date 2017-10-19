//Author : Daire Ni Chathain
//Go exercise 2 : Guessing game with h1 tag

package main

import (
	"fmt"
	"net/http"
)

func requestHandler(w http.ResponseWriter, r *http.Request) {

	//set content to html so can render <H> tags
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	//added h1 tags
	fmt.Fprintln(w, "<h1>Guessing Game</h1> ")

}

func main() {
	http.HandleFunc("/", requestHandler)
	http.ListenAndServe(":8080", nil)
}
