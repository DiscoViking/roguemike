package events

import (
	"log"
	"sync"
)

type broker struct {
	sync.RWMutex
	subscribers map[Type][]Client
	seq         int64
}

type taggedEvent struct {
	e   Event
	seq int64
}

func NewBroker() Broker {
	m := broker{}
	m.subscribers = make(map[Type][]Client, 0)
	return m
}

func (b broker) Subscribe(t Type, c Client) {
	b.Lock()
	defer b.Unlock()

	_, ok := b.subscribers[t]
	if !ok {
		b.subscribers[t] = make([]Client, 0, 1)
	}

	log.Printf("New subscriber for event type %v.", t)
	b.subscribers[t] = append(b.subscribers[t], c)
}

func (b broker) Publish(e Event) {
	b.RLock()
	defer b.RUnlock()

	tagged := b.tagEvent(e)

	t := e.Type()
	log.Printf("New event of type %v", t)
	if subscribers, ok := b.subscribers[t]; ok {
		log.Printf("Publishing %v event to %v subscribers", t, len(subscribers))
		for _, s := range subscribers {
			s.handle(tagged)
		}
	}
}

func (b broker) tagEvent(e Event) taggedEvent {
	b.seq += 1
	return taggedEvent{
		e:   e,
		seq: b.seq,
	}
}
