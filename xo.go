package xo

type Game struct {
	Players []Player
	Board   Board
}

func NewGame(playerOne, playerTwo Player, boxSize int) Game {
	size := NewSize(boxSize)
	board := NewBoard(size)
	return Game{
		Players: []Player{playerOne, playerTwo},
		Board:   board,
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

func NewSize(boxSize int) Size {
	return Size{boxSize, boxSize}
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
