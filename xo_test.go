package xo

import "testing"

func Test_NewGame_Input_2_Players_BoxSize_3_Should_Be_Game_BoxSize_3_With_2_Players(t *testing.T) {
	playerOne := NewPlayer("O")
	playerTwo := NewPlayer("X")
	boxSize := 3
	expected := Game{
		Players: []Player{
			playerOne,
			playerTwo,
		},
		Board: Board{
			Slots: [][]State{
				[]State{State{}, State{}, State{}},
				[]State{State{}, State{}, State{}},
				[]State{State{}, State{}, State{}},
			},
			Size: Size{
				X: 3,
				Y: 3,
			},
		},
	}

	actual := NewGame(playerOne, playerTwo, boxSize)

	if len(expected.Players) != len(actual.Players) ||
		expected.Board.Size != actual.Board.Size {
		t.Errorf("Expected %v but it got %v", expected, actual)
	}
}
func Test_NewGame_Input_2_Players_Size_5_Should_Be_Game_BoxSize_5_With_2_Players(t *testing.T) {
	playerOne := NewPlayer("O")
	playerTwo := NewPlayer("X")
	boxSize := 5
	expected := Game{
		Players: []Player{
			playerOne,
			playerTwo,
		},
		Board: Board{
			Slots: [][]State{
				[]State{State{}, State{}, State{}},
				[]State{State{}, State{}, State{}},
				[]State{State{}, State{}, State{}},
				[]State{State{}, State{}, State{}},
				[]State{State{}, State{}, State{}},
			},
			Size: Size{
				X: 5,
				Y: 5,
			},
		},
	}

	actual := NewGame(playerOne, playerTwo, boxSize)

	if len(expected.Players) != len(actual.Players) ||
		expected.Board.Size != actual.Board.Size {
		t.Errorf("Expected %v but it got %v", expected, actual)
	}
}

func Test_NewPlayer_Input_O_Should_Be_Player_O(t *testing.T) {
	symbol := "O"
	expected := Player{Symbol: "O"}

	actual := NewPlayer(symbol)

	if expected != actual {
		t.Errorf("Expected %v but it got %v", expected, actual)
	}
}

func Test_NewPlayer_Input_X_Should_Be_Player_X(t *testing.T) {
	symbol := "X"
	expected := Player{Symbol: "X"}

	actual := NewPlayer(symbol)

	if expected != actual {
		t.Errorf("Expected %v but it got %v", expected, actual)
	}
}

func Test_NewBoard_Players_Should_Be_Board(t *testing.T) {
	size := Size{X: 3, Y: 3}
	expected := Board{
		Slots: [][]State{
			[]State{State{}, State{}, State{}},
			[]State{State{}, State{}, State{}},
			[]State{State{}, State{}, State{}},
		},
		Size: Size{
			X: 3,
			Y: 3,
		},
	}

	actual := NewBoard(size)

	if len(expected.Slots) != len(actual.Slots) ||
		len(expected.Slots[0]) != len(actual.Slots[0]) ||
		expected.Size != actual.Size {
		t.Errorf("Expected %v but it got %v", expected, actual)
	}
}

func Test_NewSize_Input_3_Should_Be_Size_3x3(t *testing.T) {
	input := 3
	expected := Size{3, 3}

	actual := NewSize(input)
	if expected != actual {
		t.Errorf("Expected %v but it got %v", expected, actual)
	}
}

func Test_TurnOf_Input_Player_O_Position_2_2_Should_Be_Position_2_2_Symbol_O(t *testing.T) {
	playerOne := NewPlayer("O")
	playerTwo := NewPlayer("X")
	game := NewGame(playerOne, playerTwo, 3)
	expectedState := State{Symbol: "O"}

	game.TurnOf(playerOne, 2, 2)
	actual := game.GetBoardPosition(2, 2)

	if expectedState != actual {
		t.Errorf("Expected %v but it got %v", expectedState, actual)
	}

}

func Test_IsBoardEmpty_Input_2_2_Should_Be_True(t *testing.T) {
	playerOne := NewPlayer("O")
	playerTwo := NewPlayer("X")
	game := NewGame(playerOne, playerTwo, 3)
	expected := true

	actual := game.IsBoardEmpty(2, 2)

	if expected != actual {
		t.Errorf("Expected %v but it got %v", expected, actual)
	}

}

func Test_IsBoardEmpty_Input_2_2_Should_Be_False(t *testing.T) {
	playerOne := NewPlayer("O")
	playerTwo := NewPlayer("X")
	game := NewGame(playerOne, playerTwo, 3)
	expected := false

	game.TurnOf(playerOne, 2, 2)
	actual := game.IsBoardEmpty(2, 2)

	if expected != actual {
		t.Errorf("Expected %v but it got %v", expected, actual)
	}

}

func Test_Fill_Input_Position_2_2_Player_O_Should_Be_Position_2_2_Symbol_O(t *testing.T) {
	playerOne := NewPlayer("O")
	playerTwo := NewPlayer("X")
	game := NewGame(playerOne, playerTwo, 3)
	expectedState := State{Symbol: "O"}

	game.Fill(2, 2, playerOne.Symbol)
	actual := game.GetBoardPosition(2, 2)

	if expectedState != actual {
		t.Errorf("Expected %v but it got %v", expectedState, actual)
	}

}

func Test_GetWinner_Input_Player_O_Should_Be_False(t *testing.T) {
	playerOne := NewPlayer("O")
	playerTwo := NewPlayer("X")
	game := NewGame(playerOne, playerTwo, 3)
	expected := false

	game.Fill(2, 2, playerOne.Symbol)
	actual := game.GetWinner(playerOne, 2, 2)

	if expected != actual {
		t.Errorf("Expected %v but it got %v", expected, actual)
	}

}

func Test_GetWinner_Input_Player_O_Should_Be_True(t *testing.T) {
	playerOne := NewPlayer("O")
	playerTwo := NewPlayer("X")
	game := NewGame(playerOne, playerTwo, 3)
	expected := true

	game.Fill(2, 2, playerOne.Symbol)
	game.Fill(2, 1, playerOne.Symbol)
	game.Fill(2, 0, playerOne.Symbol)
	actual := game.GetWinner(playerOne, 2, 0)

	if expected != actual {
		t.Errorf("Expected %v but it got %v", expected, actual)
	}

}

func Test_GetWinner_Input_Player_X_Should_Be_True(t *testing.T) {
	playerOne := NewPlayer("O")
	playerTwo := NewPlayer("X")
	game := NewGame(playerOne, playerTwo, 3)
	expected := true

	game.Fill(2, 2, playerTwo.Symbol)
	game.Fill(1, 2, playerTwo.Symbol)
	game.Fill(0, 2, playerTwo.Symbol)
	actual := game.GetWinner(playerTwo, 0, 2)

	if expected != actual {
		t.Errorf("Expected %v but it got %v", expected, actual)
	}

}
func Test_GetWinner_Input_Player_O_Should_Be_True_Diagonal(t *testing.T) {
	playerOne := NewPlayer("O")
	playerTwo := NewPlayer("X")
	game := NewGame(playerOne, playerTwo, 3)
	expected := true

	game.Fill(0, 2, playerOne.Symbol)
	game.Fill(1, 1, playerOne.Symbol)
	game.Fill(2, 0, playerOne.Symbol)
	actual := game.GetWinner(playerOne, 2, 0)

	if expected != actual {
		t.Errorf("Expected %v but it got %v", expected, actual)
	}

}

func Test_GetWinner_Input_Player_X_Should_Be_True_Diagonal(t *testing.T) {
	playerOne := NewPlayer("O")
	playerTwo := NewPlayer("X")
	game := NewGame(playerOne, playerTwo, 3)
	expected := true

	game.Fill(0, 0, playerTwo.Symbol)
	game.Fill(1, 1, playerTwo.Symbol)
	game.Fill(2, 2, playerTwo.Symbol)
	actual := game.GetWinner(playerTwo, 2, 2)

	if expected != actual {
		t.Errorf("Expected %v but it got %v", expected, actual)
	}

}

func Test_TurnOf_Input_Player_O_Should_Be_True(t *testing.T) {
	playerOne := NewPlayer("O")
	playerTwo := NewPlayer("X")
	game := NewGame(playerOne, playerTwo, 3)
	expected := true

	actual := game.TurnOf(playerOne, 2, 2)
	if expected != actual {
		t.Errorf("Expected %v but it got %v", expected, actual)
	}
}

func Test_TurnOf_Input_Player_X_Should_Be_False(t *testing.T) {
	playerOne := NewPlayer("O")
	playerTwo := NewPlayer("X")
	game := NewGame(playerOne, playerTwo, 3)
	expected := false

	actual := game.TurnOf(playerTwo, 2, 2)
	if expected != actual {
		t.Errorf("Expected %v but it got %v", expected, actual)
	}
}

func Test_NextPlayer_Current_Player_O_Should_Be_Player_X(t *testing.T) {
	playerOne := NewPlayer("O")
	playerTwo := NewPlayer("X")
	game := NewGame(playerOne, playerTwo, 3)
	expected := playerTwo

	game.NextPlayer()
	actual := game.CurrentPlayer
	if expected != actual {
		t.Errorf("Expected %v but it got %v", expected, actual)
	}
}

func Test_NextPlayer_Current_Player_X_Should_Be_Player_O(t *testing.T) {
	playerOne := NewPlayer("O")
	playerTwo := NewPlayer("X")
	game := NewGame(playerOne, playerTwo, 3)
	expected := playerOne

	game.NextPlayer()
	game.NextPlayer()
	actual := game.CurrentPlayer
	if expected != actual {
		t.Errorf("Expected %v but it got %v", expected, actual)
	}
}

func Test_IsContinue_with_Empty_Board_Should_Be_True(t *testing.T) {
	playerOne := NewPlayer("O")
	playerTwo := NewPlayer("X")
	game := NewGame(playerOne, playerTwo, 3)
	expected := true

	actual := game.IsContinue()
	if expected != actual {
		t.Errorf("Expected %v but it got %v", expected, actual)
	}
}

func Test_IsContinue_with_Full_Board_Should_Be_False(t *testing.T) {
	playerOne := NewPlayer("O")
	playerTwo := NewPlayer("X")
	game := NewGame(playerOne, playerTwo, 3)
	expected := false

	game.Board.SlotLeft = 0
	actual := game.IsContinue()
	if expected != actual {
		t.Errorf("Expected %v but it got %v", expected, actual)
	}
}
