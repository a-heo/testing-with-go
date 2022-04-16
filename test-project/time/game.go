package poker

//decouples CLI reliance on TexasHoldem type
type Game interface {
	Start(numberOfPlayers int)
	Finish(winner string)
}