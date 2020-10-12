package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"os"
)
func contact(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","text/html")
		fmt.Fprint(w,"To get in touch,please send an email to<a href=\"mailto:example@gmail.com\"> example@gmail.com</a>")

}

func home(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","text/html")
		fmt.Fprint(w,"<h1>welcome to my awesome  home site!</h1>")
}

//start registering a couple of URL paths and handlers
func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", home)
	r.HandleFunc("/contact",contact)
	port := os.Getenv("ROUTE")     //set in Dockerfile
	http.ListenAndServe(port,r)
}

