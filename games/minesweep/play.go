package minesweep

import "github.com/evan-forbes/train/play"

type Input struct {
	X, Y int
}

func (i *Input) Class() play.MessageClass {
	return play.STANDARD
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
	return []byte{byte(m.Xlen), byte(m.Ylen)}
}

func (m *MineSweep) Play(input <-chan []byte) <-chan []byte {
	out := make(chan []byte)
	go func() {
		defer close(out)
		for in := range input {
			if len(in) > 1024 {
				// m.EndGame()
			}
			m.Moves++
			out <- m.Parse(in)
		}
	}()
	return out
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
