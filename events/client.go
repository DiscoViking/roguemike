package events

import "sync"

const (
	CLIENT_EVENT_BUFFER_SIZE = 10
)

type client struct {
	sync.RWMutex
	c        chan Event
	handlers map[Type]Handler
	broker   Broker
}

func NewClient(b Broker) Client {
	return &client{
		c:        make(chan Event, CLIENT_EVENT_BUFFER_SIZE),
		handlers: make(map[Type]Handler, 0),
		broker:   b,
	}
}

func (c *client) ListenForever() {
	for e := range c.c {
		c.handleInternal(e)
	}
}

func (c *client) Subscribe(t Type, h Handler) {
	c.Lock()
	c.handlers[t] = h
	c.Unlock()

	c.broker.Subscribe(t, c)
}

func (c *client) Publish(e Event) {
	c.broker.Publish(e)
}

func (c *client) Handle(e Event) {
	c.c <- e
}

func (c *client) handleInternal(e Event) {
	c.RLock()
	defer c.RUnlock()

	t := e.Type()

	if h, ok := c.handlers[t]; ok {
		h.Handle(e)
	}
}
