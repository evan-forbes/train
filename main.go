package main

import (
	"fmt"
	"strings"

	"github.com/evan-forbes/train/games/cards"
)

const suites = "H/D/S/C"
const values = "2/3/4/5/6/7/8/9/10/J/Q/K/A"

var players = []string{"Taco Joe", "churchill", "roosevelt"}

var allCards = cards.CombineMany(
	strings.Split(values, "/"),
	strings.Split(suites, "/"),
)

func main() {
	// cmd.Execute()
	d := cards.NewDeck(players, allCards)
	fmt.Println(d.Cards, "\n", d.Dealt)
	d.DealRounds(players, 2)
	fmt.Println(d.Cards, "\n", d.Dealt)
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
