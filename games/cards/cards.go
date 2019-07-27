package cards

import (
	"errors"
	"math/rand"
	"sync"
	"time"
)

////////////////////////////////////////
// 	  Card and hand: base data units
//////////////////////////////////////

// Card wraps string to repesent an
// immutable card in a deck
type Card string

// hand is the representation of what
// cards a player holds
type hand map[Card]int

func cardConvert(in []string) []Card {
	var out []Card
	for _, i := range in {
		out = append(out, Card(i))
	}
	return out
}

////////////////////////////////////////
// 	   	Deck: Card managment
//////////////////////////////////////

// Deck keeps track of cards, and fullfills the typical
// representation of a deck of cards, but does not remain
// ordered. Faster than an ordered deck.
type Deck struct {
	*sync.Mutex
	Cards    []Card
	Discards []Card
	Dealt    map[string]hand
	RSource  *rand.Rand
}

func NewDeck(players []string, cards []string) *Deck {
	src := rand.NewSource(time.Now().UnixNano() + rand.Int63n(100))
	hands := make(map[string]hand)
	for _, player := range players {
		h := make(hand)
		hands[player] = h
	}
	return &Deck{
		Cards:   cardConvert(cards),
		Dealt:   hands,
		RSource: rand.New(src),
	}
}

// pull draws a card from the slice
func (d *Deck) Draw(i int, s []Card) (Card, error) {
	deckSize := len(s)
	if deckSize == 0 || deckSize < i {
		return "", errors.New("Deck is too small or empty")
	}
	last := deckSize - 1
	d.Lock()
	out := s[i]
	// erase the pulled card by replacing
	// it with the last card then trunkating
	s[i] = s[last]
	s[last] = ""
	s = s[:last]
	d.Unlock()
	return out, nil
}

// DrawRando takes a card out of the deck. It does not
// maintain order of the deck, but does run O(1)
func (d *Deck) DrawRando() Card {
	cardID := d.RSource.Intn(len(d.Cards))
	out, err := d.Draw(cardID, d.Cards)
	return out
}

func (d *Deck) Deal(player string, card Card) {
	d.Lock()
	d.Dealt[player] = append(d.Dealt[player], card)
	d.Unlock()
}

func (d *Deck) DealRando(player string) {
	card := d.DrawRando()
	d.Deal()

}

func (d *Deck) DealRounds(count int) {
	for i := 0; i < count; i++ {

	}
}

////////////////////////////////////////
// 	   String Utility Functions
//////////////////////////////////////

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
