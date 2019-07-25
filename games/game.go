package games

// when creating a game, user creates a channel of game specific input type
// that follows the UserInput interface

// UserInput helps describe the structure of a game
// while leaving the struct to still be completely
// game specific
type UserInput interface {
	// Type acts as a message from the user to the game
	Type() string
}

// Feedback helps describe the structure of a game
// while leaving the struct to still be completely
// game specific
type Feedback interface {
	// Info acts as
	Info() string
}

// Game describes the way a player and a game communicate
type Game interface {
	// Play accepts the player's id as a string
	// along with a user generated channel for input
	Play(string, <-chan UserInput) <-chan Feedback
}
