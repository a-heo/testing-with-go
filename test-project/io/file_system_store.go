package main

import (
	"io"
	"encoding/json"
)

//modified to writer only instead of readwriteseeeker to reduce future bugs
type FileSystemPlayerStore struct{
	database io.Writer
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
	//able to remove seek since write encapsulates seek
	// f.database.Seek(0,0)
	json.NewEncoder(f.database).Encode(f.league)
}

//construct and store league as val into struct to be used as reads during initialization 
func NewFileSystemPlayerStore(database io.ReadWriteSeeker) *FileSystemPlayerStore {
	database.Seek(0, 0)
	league, _ := NewLeague(database)
	return &FileSystemPlayerStore{
		database: &tape{database},
		league: league,
	}
}