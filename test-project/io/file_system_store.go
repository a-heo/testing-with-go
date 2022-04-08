package main

import (
	"io"
	"encoding/json"
)

//readseeker allows us to read multiple times compared to reader(only once)
type FileSystemPlayerStore struct{
	database io.ReadWriteSeeker
}

func (f *FileSystemPlayerStore) GetLeague() []Player{
	f.database.Seek(0, 0)
	league, _ := NewLeague(f.database)
	return league
}

func (f *FileSystemPlayerStore) GetPlayerScore(name string) int {
	var wins int
	for _, player := range f.GetLeague() {
		if player.Name == name {
			wins = player.Wins
			break
		}
	}
	return wins
}

func (f *FileSystemPlayerStore) RecordWin(name string)  {
	league := f.GetLeague()

	//updating index i rather than player[win] because we're ranging over a slice so we're looping over a copy of element. need to get reference actual value to change value 
	for i, player := range league {
		if player.Name == name {
			league[i].Wins++
		}
	}
	f.database.Seek(0,0)
	json.NewEncoder(f.database).Encode(league)
}