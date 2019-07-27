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
