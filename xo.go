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

func (g Game) GetWinner(player Player, row, col int) bool {
	verticle := AllRow(g.Board.Slots[row], func(state State) bool {
		return state.Symbol == player.Symbol
	})
	horizontal := AllCol(g.Board.Slots, col, func(state State) bool {
		return state.Symbol == player.Symbol
	})
	return verticle || horizontal
}

func AllRow(vs []State, f func(State) bool) bool {
	for _, v := range vs {
		if !f(v) {
			return false
		}
	}
	return true
}
func AllCol(vs [][]State, col int, f func(State) bool) bool {
	for i, _ := range vs {
		if !f(vs[i][col]) {
			return false
		}
	}
	return true
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
