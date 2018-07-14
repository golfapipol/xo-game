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

func NewPlayer(symbol string) Player {
	return Player{
		Symbol: symbol,
	}
}

func NewBoard(size Size) Board {
	state := make([][]State, size.Y)
	for index := 0; index < size.Y; index++ {
		state[index] = make([]State, size.X)
	}
	return Board{
		Slots: state,
		Size:  size,
	}
}

type Board struct {
	Slots [][]State
	Size  Size
}

type State struct {
	Symbol string
}

type Player struct {
	Symbol string
}

type Size struct {
	X int
	Y int
}
