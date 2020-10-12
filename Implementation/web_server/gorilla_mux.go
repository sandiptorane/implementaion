package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)
 func handlerFunc(w http.ResponseWriter, r *http.Request){
 	w.Header().Set("Content-Type","text/html")
 	if r.URL.Path == "/contact"{
		fmt.Fprint(w,"To get in touch,please send an email to<a href=\"mailto:example@gmail.com\"> example@gmail.com</a>")
	}else{
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w,"<h1> we could not find the page you were looking for</h1>")
	}
 }

 func home(w http.ResponseWriter, r *http.Request){
 	w.Header().Set("Content-Type","text/html")
 	if r.URL.Path == "/"{
		fmt.Fprint(w,"<h1>welcome to my awesome  home site!</h1>")
	}else{
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w,"<h1> we could not find the page you were looking for</h1>")
	}
 }

//start registering a couple of URL paths and handlers
func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", home)
	r.HandleFunc("/contact",handlerFunc)
	http.ListenAndServe(":3000",r)
}
