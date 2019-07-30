package deck

import "strings"

// I'm manipulating []byte here manually, but you of course
// could write a more typical type that follows the io.Reader
// and io.Writer interfaces (probably be a better idea).

// 			STANDARD DECK
const standardSuites = "H/D/S/C"
const standardValues = "2/3/4/5/6/7/8/9/10/J/Q/K/A"

// addRank adds index of each string in []string
// to the front of the string
// NOTE: adding rank to the cards themselves is
// far from necessary. ranking should probably be
// handled by the game itself, not the deck.
// and this function will only work for slices less
// than 255 in len
func addRank(in []string) []string {
	var out []string
	for i, s := range in {
		var j []byte
		j = append(j, byte(i))
		j = append(j, []byte(s)...)
		out = append(out, string(j))
	}
	return out
}

// standardDeck makes a typical deck of cards with
// the a high and two low.
func StandardDeck() *Deck {
	ranked := addRank(strings.Split(standardValues, "/"))
	allCards := CombineMany(ranked, strings.Split(standardSuites, "/"))
	return New(allCards)
}

///////////////// OR //////////////////

type standCard struct {
	suit, val string
	rank      int
}

func (card *standCard) Read(p []byte) (int, error) {
	p = append(p, byte(card.rank))
	p = append(p, []byte(card.val+card.suit)...)
	return len(p), nil
}

// 			PINOCHLE DECK
const pinochleValues = "9/J/Q/K/10/A"

// pinochleDeck returns a deck made for the game
// pinochle. This deck has multiple copies of a
// single card, so it's a decent example of how
// to do that
func pinochleDeck() *Deck {
	ranked := addRank(strings.Split(pinochleValues, "/"))
	// copying each ranked card
	for _, card := range ranked {
		ranked = append(ranked, card)
	}
	// combining
	allCards := CombineMany(ranked, strings.Split(standardSuites, "/"))
	return New(allCards)
}

// 			POKEMON DECK
var pokemon = []string{
	"Bulbasaur", "Ivysaur", "Venusaur", "Charmander",
	"Charmeleon", "Charizard", "Squirtle", "Wartortle",
	"Blastoise", "Caterpie", "Metapod", "Butterfree",
	"Weedle", "Kakuna", "Beedrill", "Pidgey", "Pidgeotto",
	"Pidgeot", "Rattata", "Raticate", "Spearow", "Fearow",
	"Ekans", "Arbok", "Pikachu", "Raichu", "Sandshrew",
	"Sandslash", "Nidoran♀ ", "Nidorina ", "Nidoqueen",
	"Nidoran", "Nidorino", "Nidoking", "Clefairy", "Clefable",
	"Vulpix", "Ninetales", "Jigglypuff", "Wigglytuff",
	"Zubat", "Golbat", "Oddish", "Gloom", "Vileplume",
	"Paras", "Parasect ", "Venonat ", "Venomoth", "Diglett",
	"Dugtrio ", "Meowth", "Persian", "Psyduck", "Golduck",
	"Mankey", "Primeape", "Growlithe", "Arcanine ", "Poliwag",
	"Poliwhirl", "Poliwrath", "Abra", "Kadabra", "Alakazam",
	"Machop", "Machoke", "Machamp", "Bellsprout", "Weepinbell",
	"Victreebel", "Tentacool ", "Tentacruel", "Geodude ",
	"Graveler", "Golem ", "Ponyta", "Rapidash", "Slowpoke",
	"Slowbro", "Magnemite", "Magneton ", "Farfetch'd", "Doduo",
	"Dodrio", "Seel", "Dewgong", "Grimer", "Muk", "Shellder",
	"Cloyster", "Gastly", "Haunter", "Gengar", "Onix",
	"Drowzee", "Hypno", "Krabby", "Kingler", "Voltorb",
	"Electrode", "Exeggcute", "Exeggutor", "Cubone",
	"Marowak", "Hitmonlee", "Hitmonchan", "Lickitung",
	"Koffing", "Weezing", "Rhyhorn", "Rhydon", "Chansey",
	"Tangela", "Kangaskhan", "Horsea", "Seadra", "Goldeen",
	"Seaking", "Staryu", "Starmie", "Mr. Mime", "Scyther",
	"Jynx", "Electabuzz", "Magmar", "Pinsir", "Tauros",
	"Magikarp", "Gyarados", "Lapras", "Ditto", "Eevee", "Vaporeon",
	"Jolteon", "Flareon", "Porygon", "Omanyte", "Omastar",
	"Kabuto", "Kabutops", "Aerodactyl", "Snorlax", "Articuno",
	"Zapdos", "Moltres", "Dratini", "Dragonair", "Dragonite",
	"Mewtwo", "Mew"}

func pokemonDeck() *Deck {
	return New(addRank(pokemon))
}
