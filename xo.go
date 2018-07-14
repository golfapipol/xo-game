package xo

type Game struct {
	Players []Player
	Board   Board
}

func (g Game) TurnOf(player Player, row, col int) {
	if g.IsBoardEmpty(row, col) {
		g.Fill(row, col, player.Symbol)
	}
}

func (g Game) GetBoardPosition(row, col int) State {
	return g.Board.Slots[row][col]
}

func (g Game) IsBoardEmpty(row, col int) bool {
	return g.Board.Slots[row][col].Symbol == ""
}

func (g Game) Fill(row, col int, symbol string) {
	g.Board.Slots[row][col].Symbol = symbol
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

type Position struct {
	Row int
	Col int
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
