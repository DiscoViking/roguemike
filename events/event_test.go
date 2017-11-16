package events

import (
	"testing"
)

const (
	DUMMY_EVENT_1 Type = iota
	DUMMY_EVENT_2
)

type DummyEvent struct {
	t Type
}

func (e *DummyEvent) Type() Type {
	return e.t
}

func TestNewManager(t *testing.T) {
	_ = NewManager()
}

func TestOneEvent(t *testing.T) {
	m := NewManager()
	c := 0

	m.Subscribe(DUMMY_EVENT_1, func(e Event) { c += 1 })
	m.Publish(&DummyEvent{DUMMY_EVENT_1})

	if c != 1 {
		t.Errorf("Should have recieved one event, actually got %v", c)
	}
}

func TestWrongEventType(t *testing.T) {
	m := NewManager()
	c := 0
	m.Subscribe(DUMMY_EVENT_1, func(e Event) { c += 1 })
	m.Publish(&DummyEvent{DUMMY_EVENT_2})

	if c != 0 {
		t.Errorf("Should have recieved no events,  got %v", c)
	}
}

func TestMultipleTypes(t *testing.T) {
	m := NewManager()
	c := 0
	d := 0
	m.Subscribe(DUMMY_EVENT_1, func(e Event) { c += 1 })
	m.Subscribe(DUMMY_EVENT_2, func(e Event) { d += 1 })
	m.Publish(&DummyEvent{DUMMY_EVENT_2})

	if c != 0 {
		t.Errorf("Should have recieved no events of type myEvent, got %v", c)
	}
	if d != 1 {
		t.Errorf("Should have recieved no events of type otherEvent, got %v", d)
	}
}
