package xo

import (
	"fmt"
)

type ConsoleGame struct {
	Game
}

func NewConsoleGame() ConsoleGame {
	playerOne := NewPlayer("O")
	playerTwo := NewPlayer("X")

	return ConsoleGame{Game: NewGame(playerOne, playerTwo, 3)}
}

func (cg ConsoleGame) Start() {
	fmt.Println("XO Game")
	for cg.IsContinue() {
		fmt.Println(cg.DisplayOXBoard())
		row, col := cg.WaitForInput()
		cg.TurnOf(cg.CurrentPlayer, row, col)
		if cg.GetWinner(cg.CurrentPlayer, row, col) {
			cg.Winner = cg.CurrentPlayer
			break
		}
		cg.NextPlayer()
	}
	fmt.Println(cg.DisplayOXBoard())
	fmt.Println(cg.DisplayWinner())

}

func (cg ConsoleGame) WaitForInput() (int, int) {
	var row, col int
	fmt.Printf("Turn of Player %s\nPlease input row:", cg.CurrentPlayer.Symbol)
	fmt.Scanf("%d", &row)
	fmt.Printf("Please input col:")
	fmt.Scanf("%d", &col)
	return row, col
}

func (cg ConsoleGame) DisplayOXBoard() string {
	var board string
	for rowIndex := range cg.Board.Slots {
		for colIndex := range cg.Board.Slots[rowIndex] {
			if cg.Board.Slots[rowIndex][colIndex].Symbol == "" {
				board += " "
			}
			board += cg.Board.Slots[rowIndex][colIndex].Symbol
			if colIndex != (len(cg.Board.Slots[rowIndex]) - 1) {
				board += "|"
			}
		}
		board += "\n"
	}
	return board
}

func (cg ConsoleGame) DisplayWinner() string {
	if cg.Winner.Symbol == "" {
		return fmt.Sprint("DRAW")
	}
	return fmt.Sprint(cg.Winner.Symbol, "Win")
}
