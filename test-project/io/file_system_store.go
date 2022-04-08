package main

import (
	"io"
	"encoding/json"
)

//readseeker allows us to read multiple times compared to reader(only once)
type FileSystemPlayerStore struct{
	database io.ReadWriteSeeker
	league League
}

func (f *FileSystemPlayerStore) GetLeague() League {
	//still unsure of seek's input
	// f.database.Seek(0, 0)
	// league, _ := NewLeague(f.database)
	return f.league
}

func (f *FileSystemPlayerStore) GetPlayerScore(name string) int {
	// league := f.GetLeague()
	player := f.league.Find(name)

	if player != nil {
		return player.Wins
	}
	return 0
}

func (f *FileSystemPlayerStore) RecordWin(name string)  {
	// league := f.GetLeague()
	player := f.league.Find(name)

	if player != nil {
		player.Wins++
	} else {
		f.league = append(f.league, Player{name, 1})
	}

	f.database.Seek(0,0)
	json.NewEncoder(f.database).Encode(f.league)
}

//construct and store league as val into struct to be used as reads 
func NewFileSystemPlayerStore(database io.ReadWriteSeeker) *FileSystemPlayerStore {
	database.Seek(0, 0)
	league, _ := NewLeague(database)
	return &FileSystemPlayerStore{
		database: database,
		league: league,
	}
}