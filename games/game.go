package games

// Cannal is a typical two way in memory
// communication unit, directing msgs of []byte
// and error to prespecified handlers of each type
type Cannal struct {
	Conns       map[string]chan<- []byte
	Handlers    map[byte]Handler
	ErrHandlers map[error]ErrHandler
}

// Might change this is a func rather than an interface
type Handler interface {
	Handle(id string, msg []byte) error
}

// func (c *Cannal) AddHandler(b byte, h Handler) {
// 	c.Handlers[b] = h
// }

// Might change this is a func rather than an interface
type ErrHandler interface {
	Handle(id string, err error)
}

// func (c *Cannal) AddErrHandler(err error, eh ErrHandler) {
// 	c.ErrHandlers[err] = eh
// }

// Connect starts routing the incomming messages and stores the outgoing
// message channel to the id provided.
func (c *Cannal) Connect(id string, input <-chan []byte) <-chan []byte {
	outgoing := make(chan []byte)
	c.Conns[id] = outgoing
	// handle the incomming stream
	go c.routeMsgs(id, input)
	return outgoing
}

// Quit closes all communation channels to any outstanding
// connections
func (c *Cannal) Quit() {
	for _, comm := range c.Conns {
		close(comm)
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
		h, contains := c.Handlers[msg[0]]
		if !contains {
			continue
		}
		err := h.Handle(id, msg[1:])
		if err != nil {
			c.routeErrs(id, err)
		}
	}
	close(c.Conns[id])
}

// routeErrs connects an error and its reporter to
// the appropriate prescribed ErrHandler
func (c *Cannal) routeErrs(id string, err error) {

	c.ErrHandlers[err](id, err)
}
