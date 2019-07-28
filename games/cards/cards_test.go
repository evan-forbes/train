package cards

import (
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

func deckToHand(d []Card) hand {
	out := make(hand)
	for _, card := range d {
		out[card]++
	}
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
}

func diffCards(first, second []Card) []Card {
	// convert to mappings
	f, s := deckToHand(first), deckToHand(second)
	diffs := make(hand)
	for card, count := range f {
		scount, contains := s[card]
		if !contains {
			diffs[card] = count
		}
		if scount != count {
			diffs[card] = abs(count - scount)
		}

	}
}

func checkCards(old, new []Card, hs ...hand) error {
	// ensure cards are dealt properly
	var totalDeal []Cards
	// I swear I don't normally nest loops this much
	for h := range hs {
		for card, count := range h {
			for i := 0; i < count; i++ {
				totalDeal = append(totalDeal, card)
			}
		}
	}

}
