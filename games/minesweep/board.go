package minesweep

import (
	"errors"
	"fmt"
	"math/rand"
	"strings"
	"time"
)

// Square is the most basic modulary piece of board
type Square struct {
	Bomb      bool
	Uncovered bool
	// Value is the number of neighbors that have bomb
	Value int
}

// Byte returns the byte of s.Value for responding
// to a player after uncovering
func (s *Square) Byte() byte {
	return byte(s.Value)
}

// Render produces a string representation of the
// current board
func (s *Square) Render() string {
	if s.Uncovered {
		if s.Bomb {
			return "B"
		}
		return fmt.Sprintf("%d", s.Value)
	}
	return " "
}

// Board is the central type of all game logic
// for minesweep. Contains a 2d array of type
// Square
type Board struct {
	// Contents holds the data representation of the board
	Contents   [][]*Square
	Xlen, Ylen int
	rsource    *rand.Rand
}

// NewBoard functions as the constructor for the Board type
func NewBoard(x, y int) *Board {
	var conts [][]*Square
	for i := 0; i < y; i++ {
		var row []*Square
		for j := 0; j < x; j++ {
			row = append(row, &Square{
				Bomb:      false,
				Uncovered: false,
				Value:     0,
			})
		}
		conts = append(conts, row)
	}
	src := rand.NewSource(time.Now().UnixNano() + rand.Int63n(100))
	return &Board{
		Contents: conts,
		Xlen:     x,
		Ylen:     y,
		rsource:  rand.New(src),
	}
}

// Valid ensures that a given point in with the range of
// the board.
func (b *Board) valid(in *Input) bool {
	xval := in.X >= 0 && in.X < b.Xlen
	yval := in.Y >= 0 && in.Y < b.Ylen
	return xval && yval
}

// Uncover acts on the user's input to uncover a specific square
// on the board b, will return false if a bomb is on the board
// indicating failure
func (b *Board) Uncover(in *Input) (int, bool, error) {
	if !b.valid(in) {
		return 0, false, errors.New("out of range of board")
	}
	sq := b.Contents[in.Y][in.X]
	sq.Uncovered = true
	return sq.Value, sq.Bomb, nil
}

// UncoverAll reveals the value of each sqaure on the board
func (b *Board) UncoverAll() {
	for _, row := range b.Contents {
		for _, sq := range row {
			sq.Uncovered = true
		}
	}
}

// AddBomb will add a hidden bomb to the board, updating the
// state of the neighboring squares to reflect the presence
// of the bomb
func (b *Board) AddBomb(x, y int) bool {
	sq := b.Contents[y][x]
	// stop double bombing
	if sq.Bomb {
		return false
	}
	sq.Bomb = true
	// change the neighbors to reftlect bomb loc
	nebs := b.neighbors(y, x)
	for _, neb := range nebs {
		b.Contents[neb.Y][neb.X].Value++
	}
	return true
}

// AddRandomBomb picks a psuedo random part of the
// board and adds a bomb.
// will keep spinning forever if there are not
// enough spaces for new bombs.
func (b *Board) AddRandomBomb() {
	rc := b.randomCoord()
	success := b.AddBomb(rc.X, rc.Y)
	if !success {
		b.AddRandomBomb()
	}
}

// neighbors finds the surrounding squares that are on board
// of a give coordinate
func (b *Board) neighbors(y, x int) []coord {
	var nebs []coord
	for i := y - 1; i < y+2; i++ {
		if i >= 0 && i < b.Ylen {
			for j := x - 1; j < x+2; j++ {
				if j >= 0 && j < b.Xlen {
					nebs = append(nebs, coord{j, i})
				}
			}
		}
	}
	return nebs
}

// Render compiles a string representation the board at it's current
// state. Note: use uncover all for useful debugging.
func (b *Board) Render() string {
	var out []string
	for row := 0; row < b.Ylen; row++ {
		var col []string
		for ind := 0; ind < b.Xlen; ind++ {
			col = append(col, b.Contents[row][ind].Render())
		}
		out = append(out, strings.Join(col, "|"))
	}
	return strings.Join(out, "\n")
}

type coord struct {
	X, Y int
}

func (b *Board) randomCoord() coord {
	var out coord
	out.X = b.rsource.Intn(b.Xlen)
	out.Y = b.rsource.Intn(b.Ylen)
	return out
}
