package main

import (
	"github.com/evan-forbes/quickSweep/games/minesweep"
)

func main() {
	// cmd.Execute()
	board := minesweep.NewBoard(10, 10)

}

// func (b *Board) Contains(y, x int) bool {
// 	if x >= 0 && x < b.Xlen {
// 		if y >= 0 && y < b.Ylen {
// 			return true
// 		}
// 	}
// 	return false
// }

// func (b *Board) Iter() chan *Square {
// 	out := make(chan *Square)
// 	for y := 0; y < b.Ylen; y++ {
// 		for x := 0; x < b.Xlen; x++ {
// 			out <- &b.Contents[y][x]
// 		}
// 	}
// 	return out
// }
