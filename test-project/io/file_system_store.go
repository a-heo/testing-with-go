package main

import (
	"encoding/json"
	"os"
	"fmt"
)

//modified to point to jsonencoder created from initialization
type FileSystemPlayerStore struct{
	database *json.Encoder
	league League
}

func (f *FileSystemPlayerStore) GetLeague() League {

	return f.league
}

func (f *FileSystemPlayerStore) GetPlayerScore(name string) int {
	player := f.league.Find(name)

	if player != nil {
		return player.Wins
	}
	return 0
}

func (f *FileSystemPlayerStore) RecordWin(name string)  {
	player := f.league.Find(name)

	if player != nil {
		player.Wins++
	} else {
		f.league = append(f.league, Player{name, 1})
	}

	//intialized encoder in NFSPS so reference and use it now
	f.database.Encode(f.league)
}

//construct and store league as val into struct to be used as reads during initialization 
func NewFileSystemPlayerStore(file *os.File) (*FileSystemPlayerStore, error){
	file.Seek(0, 0)
	league, error := NewLeague(file)

	if error != nil {
		return nil, fmt.Errorf("problem loading player store from file %s, %v", file.Name(), error)
	}

	//initialize encoder in constructor and store in FSPS db type
	return &FileSystemPlayerStore{
		database: json.NewEncoder(&tape{file}),
		league: league,
	}, nil
}