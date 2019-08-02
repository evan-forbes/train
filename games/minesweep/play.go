package minesweep

import (
	"bytes"
	"fmt"

	"github.com/pkg/errors"

	"github.com/evan-forbes/train/comm"
)

// Input is a simplified input that the game can use
type Input struct {
	X, Y int
}

// sumByte converts some bytes into
// a single int
func sumByte(bs []byte) int {
	var out int
	for b := range bs {
		out = b + int(out)
	}
	return out
}

// processInput will take an input message and convert
// it to a usuable input, assuming that the msg is formatted
// correctly.
// Note: the format is incredibly arbitrary and will be changed
// to whatever has the evolutionary results
func processInput(msg []byte) ([2]int, error) {
	var out [2]int
	split := bytes.Split(msg, []byte{10, 0, 0, 0, 10})
	if len(split) < 2 {
		return out, errors.New("message could not be split")
	}
	out[0], out[1] = sumByte(split[0]), sumByte(split[1])
	return out, nil
}

// going to change this to be more complex and more open ended
// by adding byte togther to for input data
func (i *Input) Unmarshall(msg []byte) error {
	if len(msg) < 2 {
		return errors.New("message is too small")
	}
	xy, err := processInput(msg)
	i.X, i.Y = xy[0], xy[1]
	if err != nil {
		return errors.Wrap(err, "Could not unmarshal")
	}
	return nil
}

type Log struct {
	Move   int
	RawMsg []byte
	Input  Input
	User   string
}

func (l *Log) String() string {
	return fmt.Sprintf(
		"%d %s %d %d %v",
		l.Move, l.User, l.Input.X, l.Input.Y, l.RawMsg,
	)
}

// MineSweep represents a single game of minesweep
type MineSweep struct {
	*Board
	*comm.Cannal
	Moves  int
	GameID string
	Logs   []Log
}

func New(id string, bombs int, xlen, ylen int) *MineSweep {
	game := &MineSweep{
		Board:  NewBoard(xlen, ylen),
		Cannal: comm.NewCannal(id),
		Moves:  0,
		GameID: id,
	}
	game.Handlers[byte(2)] = game.InputHandler
	return game
}

func (m *MineSweep) StartMessage() []byte {
	// x y #bombs
	return []byte{byte(1), byte(m.Xlen), byte(m.Ylen)}
}

/////////////////////////
// 		Handlers
///////////////////////

func (m *MineSweep) InputHandler(id string, msg []byte) error {
	var i *Input
	err := i.Unmarshall(msg)
	if err != nil {
		newErr := errors.Wrap(err, "invalid input")
		return comm.Classify(comm.NONFATAL, newErr)
	}
	// continue to count moves
	m.Moves++
	val, bomb, err := m.Uncover(i)
	if err != nil {
		newErr := errors.Wrap(err, "invalid input")
		return comm.Classify(comm.NONFATAL, newErr)
	}
	if bomb {
		m.Quit()
		// report stats
	}
	// m.Conns[id] <- some Processing function with the byte data also provided.
	return nil
}
