package events

import (
	"log"
	"sync"
)

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
	c := &client{
		c:        make(chan Event, CLIENT_EVENT_BUFFER_SIZE),
		handlers: make(map[Type]Handler, 0),
		broker:   b,
	}
	go c.listenForever()
	return c
}

func (c *client) listenForever() {
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
	t := e.Type()
	log.Printf("Received event of type %v", t)

	c.RLock()
	h, ok := c.handlers[t]
	c.RUnlock()

	if !ok {
		log.Fatalf("Didn't have a handler for event!")
	}
	h.Handle(e)
}
