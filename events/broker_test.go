package events

import (
	"testing"
)

const (
	DUMMY_EVENT_1 Type = iota
	DUMMY_EVENT_2
)

type dummyEvent struct {
	t Type
}

func (e dummyEvent) Type() Type {
	return e.t
}

type countingHandler struct {
	count int
}

func (h *countingHandler) Handle(e Event) {
	h.count += 1
}

func TestNewBroker(t *testing.T) {
	_ = NewBroker()
}

func TestOneEvent(t *testing.T) {
	b := NewBroker()

	h := countingHandler{}
	b.Subscribe(DUMMY_EVENT_1, &h)
	b.Publish(dummyEvent{DUMMY_EVENT_1})

	if h.count != 1 {
		t.Errorf("Should have recieved one event, actually got %v", h.count)
	}
}

func TestWrongEventType(t *testing.T) {
	b := NewBroker()
	h := countingHandler{}
	b.Subscribe(DUMMY_EVENT_1, &h)
	b.Publish(dummyEvent{DUMMY_EVENT_2})

	if h.count != 0 {
		t.Errorf("Should have recieved no events,  got %v", h.count)
	}
}

func TestMultipleTypes(t *testing.T) {
	b := NewBroker()
	h1 := countingHandler{}
	h2 := countingHandler{}
	b.Subscribe(DUMMY_EVENT_1, &h1)
	b.Subscribe(DUMMY_EVENT_2, &h2)
	b.Publish(dummyEvent{DUMMY_EVENT_2})

	if h1.count != 0 {
		t.Errorf("Should have recieved no events of type myEvent, got %v", h1.count)
	}
	if h2.count != 1 {
		t.Errorf("Should have recieved no events of type otherEvent, got %v", h2.count)
	}
}
