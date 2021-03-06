//Author : Daire Ni Chathain
//Go exercise 4 : Guessing game with Bootstrap
//Resources:
// 1) http://goinbigdata.com/example-of-using-templates-in-golang/
// 2) https://stackoverflow.com/questions/26559557/how-do-you-serve-a-static-html-file-using-a-go-web-server
// 3) https://medium.com/@IndianGuru/understanding-go-s-template-package-c5307758fab0

package main

import (
    "html/template"
	"net/http"
)
//Messgage struct that hold values that will be parsed & rendered using the guess.tmpl
type message struct {
	Message string
}

// when request come in it goes to this function first
func requestHandler(w http.ResponseWriter, r *http.Request) {
	//serve the index file on load
	http.ServeFile(w, r, "index.html")
}

//the handler that gets executed when a request to /guess is made
func gameHandler(w http.ResponseWriter, r *http.Request) {
	// Set Message
	msg := message{Message: "Guess a number between 1 & 20!"}
	// Guess template file that will display the message struct
	t, _ := template.ParseFiles("guess.tmpl")  
	t.Execute(w,msg) //use the template to render the HTML
}
func main() {
	http.HandleFunc("/", requestHandler)
	//executes when /guess
	http.HandleFunc("/guess", gameHandler) 
	//Listen on port 8080
	http.ListenAndServe(":8080", nil)
}
