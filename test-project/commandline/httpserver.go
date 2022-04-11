package main

import (
	"fmt"
	"net/http"
	"strings"
	"encoding/json"
)

const jsonContentType ="application/json"

//store player score info
type PlayerStore interface {
	GetPlayerScore(name string) int
	RecordWin(name string)
	GetLeague() League
}

//store player individual data
type Player struct {
	Name string
	Wins int
}

//http interface for player info
type PlayerServer struct {
	store PlayerStore
	http.Handler	
	//instead of router *http.ServeMux embedding http.Handler methods into struct so serveHTTP method no longer needed
}

//routing creation and direction
func NewPlayerServer(store PlayerStore) *PlayerServer {
	p := new(PlayerServer)

	p.store = store

	router := http.NewServeMux()
	router.Handle("/league", http.HandlerFunc(p.handleLeagues))
	
	router.Handle("/players/", http.HandlerFunc(p.handlePlayers))
	
	p.Handler = router

	return p
}


func (p *PlayerServer) handleLeagues(w http.ResponseWriter, r *http.Request) {
	//look into this code 
	w.Header().Set("content-type", jsonContentType)
	json.NewEncoder(w).Encode(p.store.GetLeague())
}

func (p *PlayerServer) handlePlayers(w http.ResponseWriter, r *http.Request) {
	player := strings.TrimPrefix(r.URL.Path, "/players/")
	
	switch r.Method {
	case http.MethodPost:
		p.processWin(w, player)
	case http.MethodGet:
		p.showScore(w, player)
	}
}

func (p *PlayerServer) showScore(w http.ResponseWriter, player string) {
	score := p.store.GetPlayerScore(player)

	if score == 0 {
		w.WriteHeader(http.StatusNotFound)
	}
	

	fmt.Fprint(w, score)
}

func (p *PlayerServer) processWin(w http.ResponseWriter, player string) {
	p.store.RecordWin(player)
	w.WriteHeader(http.StatusAccepted)
}
