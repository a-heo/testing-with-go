package main 

//initializes empty player store
func NewInMemoryPlayerStore() *InMemoryPlayerStore {
	return &InMemoryPlayerStore{map[string]int{}}
}

//collects data about players in memory
type InMemoryPlayerStore struct {
	store map[string]int	
}

//record player's win
func (i *InMemoryPlayerStore) RecordWin(name string) {
	i.store[name]++
}

//gets player score for specific player
func (i *InMemoryPlayerStore) GetPlayerScore(name string) int {
	return i.store[name]
}

//get all the player scores
func (i *InMemoryPlayerStore) GetLeague() League {
	var league []Player
	for name, wins := range i.store {
		league = append(league, Player{name, wins})
	}
	return league
}