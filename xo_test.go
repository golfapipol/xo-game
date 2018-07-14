package xo

import "testing"

func Test_NewGame_Input_2_Players_Should_Be_Game_With_2_Players(t *testing.T) {
	playerOne := Player{Symbol: "O"}
	playerTwo := Player{Symbol: "X"}
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
