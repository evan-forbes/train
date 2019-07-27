package cards

import (
	"math/rand"
	"sync"
	"time"
)

const suites = "H/D/S/C"
const values = "2/3/4/5/6/7/8/9/10/J/Q/K/A"

// AllCombos generates a slice of strings by
// every element of the first slice with every
// element of the second
func AllCombos(first, second []string) []string {
	var total []string
	for _, i := range first {
		for _, j := range second {
			total = append(total, i+j)
		}
	}
	return total
}

// CombineMany will call AllCombos for each
// []string provided
func CombineMany(in ...[]string) []string {
	if len(in) < 2 {
		return in[0]
	}
	curr := in[0]
	for i := 0; i < len(in)-1; i++ {
		curr = AllCombos(curr, in[i+1])
	}
	return curr
}

// Card wraps string to repesent an
// immutable card in a deck
type Card string

func cardConvert(in []string) []Card {
	var out []Card
	for _, i := range in {
		out = append(out, Card(i))
	}
	return out
}

// Deck keeps track of cards, and fullfills the typical
// representation of a deck of cards, but does not remain
// ordered. Faster than an ordered deck.
type Deck struct {
	*sync.Mutex
	Cards    []Card
	Discards []Card
	Dealt    map[string][]Card
	RSource  *rand.Rand
}

func NewDeck(players []string, cards []string) *Deck {
	src := rand.NewSource(time.Now().UnixNano() + rand.Int63n(100))
	hands := make(map[string][]Card)
	for _, player := range players {
		var hand []Card
		hands[player] = hand
	}
	return &Deck{
		Cards:   cardConvert(cards),
		Dealt:   hands,
		RSource: rand.New(src),
	}
}

func Draw(cardID int64) {
	a[i] = a[len(a)-1] // Copy last element to index i.
	a[len(a)-1] = ""   // Erase last element (write zero value).
	a = a[:len(a)-1]   // Truncate slice.
}

func (d *Deck) DealRando(player string) {
	d.Lock()
	cardID := d.RSource.Int63n(len(d.Cards))
	a[i] = a[len(a)-1] // Copy last element to index i.
	a[len(a)-1] = ""   // Erase last element (write zero value).
	a = a[:len(a)-1]   // Truncate slice.

	d.Unlock()

}

func (d *Deck) DealRounds(count int) {
	for i := 0; i < count; i++ {
		for player, hand := range d.Dealt {
			hand = append(hand, d.Cards[d.RSource.Int63n(len(d.Cards))])
		}
	}
}
