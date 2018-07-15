package xo

import "testing"

func Test_DisplayOXBoard_Should_Be_Board_String(t *testing.T) {
	console := NewConsoleGame()
	expected := " | | \n | | \n | | \n"
	actual := console.DisplayOXBoard()

	if expected != actual {
		t.Errorf("Expected \n%v but it got \n%v", expected, actual)
	}
}

func Test_DisplayWinner_Should_Be_DRAW(t *testing.T) {
	console := NewConsoleGame()
	expected := "DRAW"
	actual := console.DisplayWinner()

	if expected != actual {
		t.Errorf("Expected \n%v but it got \n%v", expected, actual)
	}
}
