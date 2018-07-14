package xo

import "testing"

func Test_NewGame_Input_2_Players_Should_Be_Game_With_2_Players(t *testing.T) {
	playerOne := NewPlayer("O")
	playerTwo := NewPlayer("X")
	expected := Game{
		Players: []Player{
			playerOne,
			playerTwo,
		},
		Board: Board{},
	}

	actual := NewGame(playerOne, playerTwo)

	if len(expected.Players) != len(actual.Players) {
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
