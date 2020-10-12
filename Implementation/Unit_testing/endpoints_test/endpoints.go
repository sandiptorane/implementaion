package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func CreateEndpoints(w http.ResponseWriter,r *http.Request){
	w.WriteHeader(200)
	w.Write([]byte("created"))

}

func main(){
	router := mux.NewRouter()
	router.HandleFunc("/create",CreateEndpoints).Methods("GET")
	log.Fatal(http.ListenAndServe(":3000",router))
}
