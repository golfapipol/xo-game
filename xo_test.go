package xo

import "testing"

func Test_NewGame_Should_Be_Game(t *testing.T) {
	expected := Game{}

	actual := NewGame()

	if expected != actual {
		t.Errorf("Expected %v but it got %v", expected, actual)
	}
}
