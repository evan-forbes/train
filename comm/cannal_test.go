package comm

import (
	"testing"
	"time"
)

type examplePlayer struct {
	*Cannal
	Name    string
	lastMsg string
}

type exampleGame struct {
	*Cannal
}

func (eg *exampleGame) AddPlayer(player *examplePlayer) {
	// establish connection
	eg.Connect(player.Cannal)
}

func (ep *examplePlayer) save(id string, msg []byte) error {
	ep.lastMsg = string(msg)
	return nil
}

func TestCannal(t *testing.T) {
	// make the player and game
	game := &exampleGame{NewCannal("test game")}
	player := &examplePlayer{NewCannal("Alice"), "Alice", ""}
	// assign the method save to handle all msgs with 1 as the first byte
	player.Handlers[byte(1)] = player.save
	game.AddPlayer(player)
	// make a message
	var msg []byte
	msg = append(msg, byte(1))
	msg = append(msg, []byte("hello")...)
	// send the message from the game to the player
	game.Conns["Alice"] <- msg
	// give the goroutine time to act
	time.Sleep(time.Millisecond * 1)
	if player.lastMsg != "hello" {
		t.Error("message not transfered")
	}
}
