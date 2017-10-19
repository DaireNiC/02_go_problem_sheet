//Author : Daire Ni Chathain
//Go exercise 3 : Guessing game with Bootstrap
//Resources: https://stackoverflow.com/questions/26559557/how-do-you-serve-a-static-html-file-using-a-go-web-server

package main

import (
	"net/http"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir("./guess")))
	http.ListenAndServe(":8080", nil)
}
