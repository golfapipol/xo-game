package xo

type Game struct {
	Players []Player
	Board   Board
}

func NewGame(playerOne, playerTwo Player) Game {
	return Game{
		Players: []Player{playerOne, playerTwo},
	}
}

type Board struct {
	Slots [][]State
}

type State struct {
	Symbol string
}

type Player struct {
	Symbol string
}
