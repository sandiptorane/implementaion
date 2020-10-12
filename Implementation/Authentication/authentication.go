package main

import (
	"fmt"
	"log"
	"net/http"
)
const
(
	CONN_HOST = "localhost"
	CONN_PORT = "8080"
	ADMIN_USER = "admin"
	ADMIN_PASSWORD = "admin"
)
func helloWorld(w http.ResponseWriter, r *http.Request) {
fmt.Fprintf(w, "Hello World!")
}
func BasicAuthentication(handler http.HandlerFunc, realm string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request){
		user, pass, ok := r.BasicAuth()
		if !ok || user!=ADMIN_USER || pass != ADMIN_PASSWORD{
			w.Header().Set("WWW-Authenticate", `Basic realm="`+realm+`"`)
			w.WriteHeader(401)
			w.Write([]byte("You are Unauthorized to access the application.\n"))
			return
		}
		handler(w, r)
	}
}
func main() {
	http.HandleFunc("/", BasicAuthentication(helloWorld, "Please enter your username and password"))
	err := http.ListenAndServe(CONN_HOST+":"+CONN_PORT, nil)
	if err != nil {
	log.Fatal("error starting http server : ", err)
	return
	}
}

