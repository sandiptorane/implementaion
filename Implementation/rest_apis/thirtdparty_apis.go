package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// Todo struct to map the response
type Todo struct {
	UserID    int    `json:"userId"`
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func main() {
	get()
	post()
	put()
	delete()
}

func get() {
	fmt.Println("1. Performing Http Get...")
	resp, err := http.Get("https://jsonplaceholder.typicode.com/todos/1")
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(resp.Body)    //read response body and return in bytes

	// Convert response body to string
	bodyString := string(bodyBytes)
	fmt.Println("API Response as String:\n",bodyString)

	// Convert response body to Todo struct
	var todoStruct Todo
	json.Unmarshal(bodyBytes, &todoStruct)
	fmt.Printf("API Response as struct:\n%+v\n",todoStruct)
}

func post() {
	fmt.Println("2. Performing Http Post...")
	todo := Todo{1, 3, "My first book", true}
	jsonReq, err := json.Marshal(todo)
	resp, err := http.Post("https://jsonplaceholder.typicode.com/todos", "application/json; charset=utf-8", bytes.NewBuffer(jsonReq))
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(resp.Body)   //read response body and return in bytes
	// Convert response body to string
	bodyString := string(bodyBytes)
	fmt.Println("API Response as String:\n",bodyString)

	// Convert response body to Todo struct
	var todoStruct Todo
	json.Unmarshal(bodyBytes, &todoStruct)
	fmt.Printf("API Response as struct:\n%+v\n", todoStruct)
}

func put() {
	fmt.Println("3. Performing Http Put...")
	todo := Todo{1, 2, "My new book", true}
	jsonReq, err := json.Marshal(todo)
	req, err := http.NewRequest(http.MethodPut, "https://jsonplaceholder.typicode.com/todos/1", bytes.NewBuffer(jsonReq))
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(resp.Body)

	// Convert response body to string
	bodyString := string(bodyBytes)
	fmt.Println("API Response as String:\n" + bodyString)

	// Convert response body to Todo struct
	var todoStruct Todo
	json.Unmarshal(bodyBytes, &todoStruct)
	fmt.Printf("API Response as struct:\n%+v\n", todoStruct)
}

func delete() {
	fmt.Println("4. Performing Http Delete...")
	todo := Todo{1, 2, "My new book", true}
	jsonReq, err := json.Marshal(todo)
	req, err := http.NewRequest(http.MethodDelete, "https://jsonplaceholder.typicode.com/todos/1", bytes.NewBuffer(jsonReq))
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(resp.Body)

	// Convert response body to string
	bodyString := string(bodyBytes)
	fmt.Println("API Response as String:\n",bodyString)
}
