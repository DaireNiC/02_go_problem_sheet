//Author : Daire Ni Chathain
//Go exercise 4 : Guessing game with Bootstrap
//Resources:
// 1) http://goinbigdata.com/example-of-using-templates-in-golang/
// 2) https://stackoverflow.com/questions/26559557/how-do-you-serve-a-static-html-file-using-a-go-web-server
// 3) https://medium.com/@IndianGuru/understanding-go-s-template-package-c5307758fab0

package main

import (
	//"fmt"
     "html/template"
	"net/http"


)

type message struct {
	Message string
//	Header string
}

// the handler function that gets executed every time a request arrives at the root
func requestHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "index.html")
}

//the handler that gets executed when a request to /guess is made
func gameHandler(w http.ResponseWriter, r *http.Request) {
	// the header 
	// Set Message
	msg := message{Message: "im the message "}
	t, _ := template.ParseFiles("guess.tmpl")
	t.Execute(w,msg)
}
func main() {
	http.HandleFunc("/", requestHandler)
	http.HandleFunc("/guess", gameHandler) //executes when /guess
	http.ListenAndServe(":8080", nil)
}