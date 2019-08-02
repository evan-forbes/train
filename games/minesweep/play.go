package minesweep

import "errors"

type Input struct {
	X, Y int
}

func (i *Input) Unmarshall(msg []byte) error {
	if len(msg) < 2 {
		return errors.New("*Failure* message is too small")
	}
}

type MineSweep struct {
	*Board
	Moves    int
	Bombs    int
	PlayerID string
	GameID   string
}

func (m *MineSweep) StartMessage() []byte {
	// x y #bombs
	return []byte{byte(1), byte(m.Xlen), byte(m.Ylen)}
}

// Note: Parsing and creating messaging is a huge thing to change
// and changing can actually be useful for building an ai,
func (m *MineSweep) Parse(in []byte) []byte {
	if len(in) == 0 {
		// m.EndGame()
	}
	if in[0] == 0 {
		// m.EndGame()
	}
	return []byte("not done")
}
