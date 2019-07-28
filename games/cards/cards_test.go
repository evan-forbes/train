package cards

import (
	"fmt"
	"reflect"
	"strings"
	"testing"
)

func TestAllCombos(t *testing.T) {
	result := AllCombos(
		[]string{"A", "B", "C"},
		[]string{"D", "E", "F"},
	)
	expected := []string{
		"AD", "AE", "AF",
		"BD", "BE", "BF",
		"CD", "CE", "CF",
	}
	if len(result) != len(expected) {
		t.Errorf(
			"unexpected number of combinations %d - %d",
			len(result), len(expected),
		)
	}

	for i := 0; i < len(result); i++ {
		if result[i] != expected[i] {
			t.Errorf("Unexpected order, got %s\nwanted %s", result, expected)
		}
	}

}

func TestCombineMany(t *testing.T) {
	nucs := []string{"A", "T", "C", "G"}
	result := CombineMany(nucs, nucs, nucs)
	expected := [64]string{
		"AAA", "AAT", "AAC", "AAG", "ATA",
		"ATT", "ATC", "ATG", "ACA", "ACT",
		"ACC", "ACG", "AGA", "AGT", "AGC",
		"AGG", "TAA", "TAT", "TAC", "TAG",
		"TTA", "TTT", "TTC", "TTG", "TCA",
		"TCT", "TCC", "TCG", "TGA", "TGT",
		"TGC", "TGG", "CAA", "CAT", "CAC",
		"CAG", "CTA", "CTT", "CTC", "CTG",
		"CCA", "CCT", "CCC", "CCG", "CGA",
		"CGT", "CGC", "CGG", "GAA", "GAT",
		"GAC", "GAG", "GTA", "GTT", "GTC",
		"GTG", "GCA", "GCT", "GCC", "GCG",
		"GGA", "GGT", "GGC", "GGG",
	}

	if len(result) != len(expected) {
		t.Errorf(
			"unexpected number of combinations %d - %d",
			len(result), len(expected),
		)
	}

	for i := 0; i < len(result); i++ {
		if result[i] != expected[i] {
			t.Errorf("Unexpected order, got %s\nwanted %s", result, expected)
		}
	}
}

///////////// Deck /////////////
func deckToHand(d []Card) hand {
	out := make(hand)
	for _, card := range d {
		out[card]++
	}
	fmt.Println("deck after converting: ", out)
	return out
}

// abs returns the absolute value of x.
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func combineHands(f, s hand) hand {
	for card, count := range s {
		for i := 0; i < count; i++ {
			f[card]++
		}
	}
	return f
}

func checkForDiffs(f, s hand) hand {
	diffs := make(hand)
	for card, count := range f {
		if card == "" {
			continue
		}
		scount, contains := s[card]
		if !contains {
			fmt.Println("diff card", card)
			diffs[card] = count
		}
		if contains && scount != count {
			fmt.Println("diff count: ", card, count, scount)
			diffs[card] = abs(count - scount)
		}
	}
	return diffs
}

func diffCards(first, second []Card) hand {
	// convert to mappings
	f, s := deckToHand(first), deckToHand(second)
	fToSDiff := checkForDiffs(f, s)
	sToFDiff := checkForDiffs(s, f)
	return combineHands(fToSDiff, sToFDiff)

}

// checkCards will
func testCards(old, new []Card, hs []hand) bool {
	diffs := diffCards(old, new)
	totalDealt := make(hand)
	for _, h := range hs {
		combineHands(totalDealt, h)
	}
	fmt.Println("diffs: ", diffs)
	fmt.Println("dealt: ", totalDealt)
	return reflect.DeepEqual(diffs, totalDealt)
}

func TestDeck(t *testing.T) {
	const suites = "H/D/S/C"
	const values = "2/3/4/5/6/7/8/9/10/J/Q/K/A"
	var players = []string{"Taco Joe", "churchill", "roosevelt"}
	var allCards = CombineMany(
		strings.Split(values, "/"),
		strings.Split(suites, "/"),
	)
	d := NewDeck(players, allCards)
	og := NewDeck(players, allCards)

	d.DealRounds(players, 4)
	fmt.Println(og.Cards)
	fmt.Println(d.Cards)
	var hands []hand
	for _, h := range d.Dealt {
		hands = append(hands, h)
	}
	ok := testCards(og.Cards, d.Cards, hands)
	if !ok {
		t.Error("Cards are missing or inflated!")
	}
}
