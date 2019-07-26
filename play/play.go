package play

// The following types are mostly a blueprint for
// communications between players and games

// MessageClass helps define what a Message contains.
type MessageClass int

const (
	START MessageClass = 1 + iota
	END
	STANDARD
	WIN
	LOSE
)

const (
	SIZELIMIT = 1024
)

// UserInput helps describe the structure of a game
// while leaving the struct to still be completely
// game specific.
type Message interface {
	// Type acts as a message from the user to the game
	Class() MessageClass
}

// Game describes the way a player and a game communicate
type Game interface {
	// Play accepts the player's id as a string
	// along with a user generated channel for input
	Play(<-chan []byte) <-chan []byte
}
