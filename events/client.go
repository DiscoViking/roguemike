package events

import "sync"

const (
	CLIENT_EVENT_BUFFER_SIZE = 10
)

type Client struct {
	sync.RWMutex
	c        chan Event
	handlers map[Type]Handler
	broker   Broker
}

type HandleFunc func(e Event)

func (f HandleFunc) Handle(e Event) {
	f(e)
}

func NewClient(b Broker) *Client {
	return &Client{
		c:        make(chan Event, CLIENT_EVENT_BUFFER_SIZE),
		handlers: make(map[Type]Handler, 0),
		broker:   b,
	}
}

func (c *Client) ListenForever() {
	for e := range c.c {
		c.handleInternal(e)
	}
}

func (c *Client) Subscribe(t Type, h Handler) {
	c.Lock()
	c.handlers[t] = h
	c.Unlock()

	c.broker.Subscribe(t, c)
}

func (c *Client) Emit(e Event) {
	c.broker.Publish(e)
}

func (c *Client) Handle(e Event) {
	c.c <- e
}

func (c *Client) handleInternal(e Event) {
	c.RLock()
	defer c.RUnlock()

	t := e.Type()

	if h, ok := c.handlers[t]; ok {
		h.Handle(e)
	}
}
