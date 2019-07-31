package deck

import (
	"errors"
	"fmt"
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

// hand is a counter for a grouping of cards.
// the typical representation of what cards a
// player holds.
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
// ordered. Faster than an ordered deck. Psuedo random.
type Deck struct {
	Mut      *sync.Mutex
	Cards    []Card
	Discards []Card
	Dealt    map[string]hand
	RSource  *rand.Rand
}

// NewDeck returns a ready to use deck
func New(cards []string) *Deck {
	src := rand.NewSource(time.Now().UnixNano() + rand.Int63n(100))
	hands := make(map[string]hand)

	return &Deck{
		Mut:     &sync.Mutex{},
		Cards:   cardConvert(cards),
		Dealt:   hands,
		RSource: rand.New(src),
	}
}

func (d *Deck) AddHand(players ...string) {
	for _, player := range players {
		h := make(hand)
		d.Dealt[player] = h
	}
}

// Draw removes and returns the card at index i
// from the deck
func (d *Deck) Draw(i int) (Card, error) {
	if len(d.Cards) == 0 || len(d.Cards) < i {
		return "", errors.New("Deck is too small or empty")
	}
	last := len(d.Cards) - 1
	d.Mut.Lock()
	out := d.Cards[i]
	// erase the pulled card by replacing
	// it with the last card then trunkating
	d.Cards[i] = d.Cards[last]
	d.Cards[last] = ""
	d.Cards = d.Cards[:last]
	d.Mut.Unlock()
	return out, nil
}

// DrawRando takes a card out of the deck. It does not
// maintain order of the deck, but does run O(1)
func (d *Deck) DrawRando() Card {
	cardID := d.RSource.Intn(len(d.Cards))
	// we can ignore this error, just checked for length
	// which is the only trigger for the error
	out, _ := d.Draw(cardID)
	return out
}

// Deal removes a card from the deck and assigns
// it to a hand
func (d *Deck) Deal(player string, card Card) {
	d.Mut.Lock()
	d.Dealt[player][card]++
	d.Mut.Unlock()
}

// DealRando takes a random card from the
// deck and puts it in a player's hand
func (d *Deck) DealRando(player string) {
	card := d.DrawRando()
	d.Deal(player, card)
}

// DealRounds repeatedly gives each hand
// the same amount of cards randomly drawn
// from the deck
func (d *Deck) DealRounds(players []string, count int) {
	for i := 0; i < count; i++ {
		for _, player := range players {
			d.DealRando(player)
		}
	}
}

// Discard removes a card from a hand and into the Discards
// slice, if the hand exists, and if the hand has a card to
// discard
func (d *Deck) Discard(player string, card Card) error {
	cardCount, contains := d.Dealt[player][card]
	if !contains {
		return fmt.Errorf(
			"Could not find hand/card combo: %s %s", player, card,
		)
	}
	if cardCount == 0 {
		return fmt.Errorf("hand does not contain card")
	}
	d.Dealt[player][card]--
	d.Mut.Lock()
	d.Discards = append(d.Discards, card)
	d.Mut.Unlock()
	return nil
}

////////////////////////////////////////
// 	   String Utility Functions
//////////////////////////////////////

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

// CombineMany produces all combinations of provided
// strings while preserving the digit index order
// [A, B] + [C, D] -> [AC, AD, BC, BD]
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
