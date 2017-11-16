package events

import (
	"log"
	"sync"
)

type broker struct {
	sync.RWMutex
	subscribers map[Type][]Handler
}

func NewBroker() Broker {
	m := broker{}
	m.subscribers = make(map[Type][]Handler, 0)
	return m
}

func (b broker) Subscribe(t Type, h Handler) {
	b.Lock()
	defer b.Unlock()

	_, ok := b.subscribers[t]
	if !ok {
		b.subscribers[t] = make([]Handler, 0, 1)
	}

	log.Printf("New subsriber for event type %v.", t)
	b.subscribers[t] = append(b.subscribers[t], h)
}

func (b broker) Publish(e Event) {
	b.RLock()
	defer b.RUnlock()

	t := e.Type()
	log.Printf("New event of type %v", t)
	if subscribers, ok := b.subscribers[t]; ok {
		log.Printf("Publishing %v event to %v subscribers", t, len(subscribers))
		for _, s := range subscribers {
			s.Handle(e)
		}
	}
}
