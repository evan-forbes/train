package games

type Game struct {
	Players  map[string]chan<- []byte
	Handlers map[byte]Handler
}

type Handler interface {
	Handle(player string, msg []byte) error
}

func (g *Game) ConnPlayer(name string, input <-chan []byte) <-chan []byte {
	outgoing := make(chan []byte)
	g.Players[name] = outgoing
	// handle the incomming stream
	return outgoing
}

// Quit closes all communation channels to players
func (g *Game) Quit() {
	for player, comm := range g.Players {
		close(comm)
	}
}

// RouteMsg takes all incoming messages from a player
//  and directs them to a correct handler. Meant to
// be used in a goroutine
func (g *Game) routeMsg(player string, msgs <-chan []byte) {
	for msg := range msgs {
		if len(msg) == 0 {
			continue
		}

	}
}
