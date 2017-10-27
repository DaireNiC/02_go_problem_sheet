//Author : Daire Ni Chathain
//Go exercise 6 : Guessing game with user input of guess
//Resources:
// 1) http://goinbigdata.com/example-of-using-templates-in-golang/
// 2) https://stackoverflow.com/questions/26559557/how-do-you-serve-a-static-html-file-using-a-go-web-server
// 3) https://medium.com/@IndianGuru/understanding-go-s-template-package-c5307758fab0
// 4) // https://golang.org/pkg/strconv/

package main

import (
    "html/template"
	"net/http"
	"strconv"
	"math/rand"
	"time"
	"fmt"
)

type message struct {
	Message string
	Guess int
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
				cookie = &http.Cookie{Name: "target",
				Value:   strconv.Itoa((rand.Intn(19) + 1)), 
				Expires: time.Now().Add(72 * time.Hour)}
		}

	// Take in user input from form
	//convert from string to int 
	guess, _ := strconv.Atoi(r.FormValue("guess"))

	// Set Message & guess
	msg := &message{Message: "Guess a number between 1 & 20!", Guess: guess}
	target, _ := strconv.Atoi(cookie.Value)

	//displays the random value for the cookie
	fmt.Println("The cookie values is: ",target)
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
