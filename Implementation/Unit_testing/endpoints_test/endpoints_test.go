//reference := www.thepolyglotdeveloper.com
package main

import (
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

type Person struct{
	Firstname string
	Lastname string
}
func router() *mux.Router{  //initialize the router
	r := mux.NewRouter()
	r.HandleFunc("/create",CreateEndpoints).Methods("GET")
	return r
}

func TestCreateEndpoints(t *testing.T) {
	response := httptest.NewRecorder()   //to record the response
	request ,_:=http.NewRequest("GET","/create",nil)
	router().ServeHTTP(response,request)
	assert.Equal(t,200,response.Code,"ok response expected")
}



