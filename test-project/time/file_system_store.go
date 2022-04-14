package poker

import (
	"encoding/json"
	"os"
	"fmt"
	"sort"
)

//stores players in file system
type FileSystemPlayerStore struct{
	database *json.Encoder
	league League
}

//func that encapsulates opening a file from path and return playerstore
func FileSystemPlayerStoreFromFile(path string) (*FileSystemPlayerStore, func(), error) {
	db, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE, 0666)

	if err != nil {
		return nil, nil, fmt.Errorf("problem opening %s %v", path, err)
	}

	closeFunc := func() {
		db.Close()
	}

	store, err := NewFileSystemPlayerStore(db)

	if err != nil {
		return nil, nil, fmt.Errorf("problem creating file system player store, %v", err)
	}

	return store, closeFunc, nil
}

//construct and store league as val into struct to be used as reads during initialization 
func NewFileSystemPlayerStore(file *os.File) (*FileSystemPlayerStore, error){
	
	err := initialisePLayerDBFile(file)

	if err != nil {
		return nil, fmt.Errorf("problem initialising playder db file, %v", err)
	}

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

//separate func from NFSPS to only initialise the file
func initialisePLayerDBFile(file *os.File) error {
	file.Seek(0, 0)
	//returns stats on file 
	info, err := file.Stat()
	if err != nil {
		return fmt.Errorf("problem getting file info from file %s, %v", file.Name(), err)
	}
	//if size is empty write an empty json array and seek to start again
	if info.Size() == 0 {
		file.Write([]byte("[]"))
		file.Seek(0,0)
	}
	
	return nil
}

//displays and sorts players in order of win
func (f *FileSystemPlayerStore) GetLeague() League {
	sort.Slice(f.league, func(i, j int) bool {
		return f.league[i].Wins > f.league[j].Wins
	})
	return f.league
}

//retrieves player score
func (f *FileSystemPlayerStore) GetPlayerScore(name string) int {
	player := f.league.Find(name)

	if player != nil {
		return player.Wins
	}
	return 0
}

//records player win
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
