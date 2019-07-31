package cards

import (
	"github.com/evan-forbes/train/deck"
	"github.com/evan-forbes/train/games"
)

type CardGame struct {
	*deck.Deck
	*games.Cannal
}
