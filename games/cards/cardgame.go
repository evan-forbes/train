package cards

import "github.com/evan-forbes/train/deck"

type CardGame struct {
	*deck.Deck
	*games.Game
}
