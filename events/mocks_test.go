package events

import (
	"reflect"
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

type testBroker struct {
	lastEvent            Event
	lastSubscribeType    Type
	lastSubscribeHandler Handler
}

func (b *testBroker) Subscribe(t Type, h Handler) {
	b.lastSubscribeType = t
	b.lastSubscribeHandler = h
}

func (b *testBroker) Publish(e Event) {
	b.lastEvent = e
}

func (b *testBroker) verifySubscribe(t *testing.T, typ Type, h Handler) {
	if b.lastSubscribeType != typ {
		t.Errorf("Client didn't register correct type with broker (was %#v", b.lastSubscribeType)
	}

	if b.lastSubscribeHandler != h {
		t.Errorf("Client didn't register itself with broker (was %#v)", b.lastSubscribeHandler)
	}
}

func (b *testBroker) verifyPublish(t *testing.T, e Event) {
	if !reflect.DeepEqual(e, b.lastEvent) {
		t.Errorf("Published event didn't match expected (wanted %#v, got %#v)", e, b.lastEvent)
	}
}
