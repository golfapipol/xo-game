package xo

type Game struct {
	Winner        Player
	Players       []Player
	CurrentPlayer Player
	Board         Board
}

func (g *Game) TurnOf(player Player, row, col int) bool {
	if g.CurrentPlayer != player {
		return false
	}
	if g.IsBoardEmpty(row, col) {
		g.Fill(row, col, player.Symbol)
		return true
	}
	return false
}

func (g *Game) IsContinue() bool {
	return g.Board.SlotLeft > 0 && g.Winner.Symbol == ""
}
func (g *Game) NextPlayer() {
	for index := range g.Players {
		if g.CurrentPlayer == g.Players[index] {
			if index == len(g.Players)-1 {
				g.CurrentPlayer = g.Players[0]
				return
			} else {
				g.CurrentPlayer = g.Players[index+1]
				return
			}
		}
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
	g.Board.SlotLeft--
}

func (g Game) GetWinner(player Player, row, col int) bool {
	verticle := AllRow(g.Board.Slots[row], player, MatchSymbol)
	horizontal := AllCol(g.Board.Slots, player, col, MatchSymbol)
	diagonalLeft := AllDiagonalLeft(g.Board.Slots, player, MatchSymbol)
	diagonalRight := AllDiagonalRight(g.Board.Slots, player, MatchSymbol)
	return verticle || horizontal || diagonalLeft || diagonalRight
}

func MatchSymbol(state State, player Player) bool {
	return state.Symbol == player.Symbol
}

func AllRow(vs []State, player Player, f func(State, Player) bool) bool {
	for _, v := range vs {
		if !f(v, player) {
			return false
		}
	}
	return true
}
func AllCol(vs [][]State, player Player, col int, f func(State, Player) bool) bool {
	for index := range vs {
		if !f(vs[index][col], player) {
			return false
		}
	}
	return true
}

func AllDiagonalLeft(vs [][]State, player Player, f func(State, Player) bool) bool {
	maxRow := len(vs)
	maxCol := len(vs[0])

	for i, j := 0, 0; i < maxRow && j < maxCol; i, j = i+1, j+1 {
		if !f(vs[i][j], player) {
			return false
		}
	}
	return true

}

func AllDiagonalRight(vs [][]State, player Player, f func(State, Player) bool) bool {
	maxRow := len(vs)
	maxCol := len(vs[0])
	lastCol := maxCol - 1
	for i, j := 0, lastCol; i < maxRow && j >= 0; i, j = i+1, j-1 {
		if !f(vs[i][j], player) {
			return false
		}
	}
	return true
}

type Board struct {
	Slots    [][]State
	Size     Size
	SlotLeft int
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
		Players:       []Player{playerOne, playerTwo},
		CurrentPlayer: playerOne,
		Board:         board,
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
		Slots:    state,
		Size:     size,
		SlotLeft: size.X * size.Y,
	}
}

func NewSize(boxSize int) Size {
	return Size{boxSize, boxSize}
}
