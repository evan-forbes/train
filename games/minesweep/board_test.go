package minesweep

import (
	"testing"
)

// TestRender tests the overall function of the minesweep package
// by rendering an uncovered board with bombs being added
func TestRender(t *testing.T) {
	board := NewBoard(10, 10)
	board.AddBomb(4, 3)
	board.AddBomb(5, 3)
	board.AddBomb(5, 3)
	board.AddBomb(6, 5)
	board.AddBomb(0, 0)
	board.UncoverAll()
	result := board.Render()
	expected := `B|1|0|0|0|0|0|0|0|0
1|1|0|0|0|0|0|0|0|0
0|0|0|1|2|2|1|0|0|0
0|0|0|1|B|B|1|0|0|0
0|0|0|1|2|3|2|1|0|0
0|0|0|0|0|1|B|1|0|0
0|0|0|0|0|1|1|1|0|0
0|0|0|0|0|0|0|0|0|0
0|0|0|0|0|0|0|0|0|0
0|0|0|0|0|0|0|0|0|0`
	if expected != result {
		t.Errorf("Problem Redering got:\n%s\nwanted\n%s", result, expected)
	}
}
