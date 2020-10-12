//reference: https://quii.gitbook.io/
package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

type PlayerStore interface {
	GetPlayerScore(name string) int
	RecordWins(name string)
}

type PlayerServer struct {
	store PlayerStore
}

type InMemoryPlayerStore struct{
	store map[string]int
}

func NewInMemoryPlayerStore() *InMemoryPlayerStore {
	return &InMemoryPlayerStore{map[string]int{}}
}


func (i *InMemoryPlayerStore) GetPlayerScore(name string) int {
	return i.store[name]
}

func (i *InMemoryPlayerStore) RecordWins(name string) {
	i.store[name]++
}


func (p *PlayerServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	player := strings.TrimPrefix(r.URL.Path, "/players/")
	switch r.Method{
	case "POST":
		p.processWins(w,player)
	case "GET":
		p.showScore(w,player)
	}

}

func (p *PlayerServer)processWins(w http.ResponseWriter,player string){
	p.store.RecordWins(player)

	w.WriteHeader(http.StatusAccepted)
}

func (p *PlayerServer)showScore(w http.ResponseWriter,player string){
	score := p.store.GetPlayerScore(player)
	if score==0{
		w.WriteHeader(http.StatusNotFound)
	}
	fmt.Fprint(w,score)
}

func main() {
	server := &PlayerServer{NewInMemoryPlayerStore()}
	if err := http.ListenAndServe(":3000", server); err != nil {
		log.Fatalf("could not listen on port 3000 %v", err)
	}
}