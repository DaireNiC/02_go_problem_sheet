//Author : Daire Ni Chathain
//Go exercise 5 : Guessing game with Bootstrap + check for cookie
//Resources:
// 1) http://goinbigdata.com/example-of-using-templates-in-golang/
// 2) https://stackoverflow.com/questions/26559557/how-do-you-serve-a-static-html-file-using-a-go-web-server
// 3) https://medium.com/@IndianGuru/understanding-go-s-template-package-c5307758fab0

package main

import (
    "html/template"
	"net/http"
	"strconv"
	"math/rand"
	"time"
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

		//random num generator
		rand.Seed(time.Now().UTC().UnixNano())

		//First check if a cookie exists already * name it target
		cookie, err := r.Cookie("target")
		if err != nil {
			//if no cookie, set one with random num between 1 and 20
			http.SetCookie(w, &http.Cookie{Name: "target", Value: strconv.Itoa(rand.Intn(19) + 1)})
			
		}

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
