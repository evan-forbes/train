package minesweep

import (
	"github.com/pkg/errors"

	"github.com/evan-forbes/train/comm"
)

type Input struct {
	X, Y int
}

// going to change this to be more complex and more open ended
// by adding byte togther to for input data
func (i *Input) Unmarshall(msg []byte) error {
	if len(msg) < 2 {
		return errors.New("message is too small")
	}
	i.X = int(msg[0])
	i.Y = int(msg[1])
	return nil
}

type MineSweep struct {
	*Board
	*comm.Cannal
	Moves  int
	GameID string
}

func New(id string, bombs int, xlen, ylen int) *MineSweep {
	game := &MineSweep{
		Board:  NewBoard(xlen, ylen),
		Cannal: comm.NewCannal(id),
		GameID: id,
	}
	game.Handlers[byte(2)] = game.InputHandler
	return game
}

func (m *MineSweep) InputHandler(id string, msg []byte) error {
	var i *Input
	err := i.Unmarshall(msg)
	if err != nil {
		newErr := errors.Wrap(err, "invalid input")
		return comm.Classify(comm.NONFATAL, err)
	}
	val, bomb, err := m.Uncover(i)
	if err != nil {
		newErr := errors.Wrap(err, "invalid input")
		return comm.Classify(comm.NONFATAL, err)
	}
	if bomb {
		m.Quit()

	}
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
