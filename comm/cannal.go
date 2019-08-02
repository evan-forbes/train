package comm

import (
	"fmt"
)

type Handler func(id string, msg []byte) error

type ErrHandler func(id string, err error)

// Cannal is a typical two way in memory
// communication unit, directing msgs of []byte
// and error to prespecified handlers of each type
type Cannal struct {
	ID          string
	Conns       map[string]chan<- []byte
	Handlers    map[byte]Handler
	ErrHandlers map[string]ErrHandler
}

func NewCannal(id string) *Cannal {
	return &Cannal{
		ID:          id,
		Conns:       make(map[string]chan<- []byte),
		Handlers:    make(map[byte]Handler),
		ErrHandlers: make(map[string]ErrHandler),
	}
}

// Connect established connections between the two cannals
// by storing channels to their repsective ids and starting
// routing
func (a *Cannal) Connect(b *Cannal) {
	go a.routeMsgs(b.ID, b.dial(a.ID))
	go b.routeMsgs(a.ID, a.dial(b.ID))
}

// dial is the first step in establishing a connection to
// another Cannal. the returned outgoing chan is stored
// in the Conns using the id provided
func (c *Cannal) dial(id string) chan []byte {
	outgoing := make(chan []byte)
	c.Conns[id] = outgoing
	return outgoing
}

// Quit closes all communation channels to any outstanding
// connections
func (c *Cannal) Quit() {
	for _, conn := range c.Conns {
		close(conn)
	}
}

// routeMsgs takes all incoming messages from an id
// and directs them to a correct handler. Meant to be used
//  in a goroutine. will block while handling message
func (c *Cannal) routeMsgs(id string, msgs <-chan []byte) {
	for msg := range msgs {
		if len(msg) == 0 {
			continue
		}
		handleFunc, contains := c.Handlers[msg[0]]
		if !contains {
			continue
		}
		err := handleFunc(id, msg[1:])
		if err != nil {
			c.routeErr(id, err)
		}
	}
	close(c.Conns[id])
}

// routeErrs connects an error and its reporter to
// the appropriate prescribed ErrHandler
func (c *Cannal) routeErr(id string, err error) {
	classyerr, ok := err.(ClassyErr)
	if !ok {
		fmt.Println("Could not handle err:", err)
		return
	}
	handleFunc, contains := c.ErrHandlers[classyerr.Class()]
	if !contains {
		return
	}
	handleFunc(id, err)
}
