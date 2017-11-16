package events

import (
	"testing"
)

func TestRegisterHandler(t *testing.T) {
	b := &testBroker{}
	c := NewClient(b)

	h := countingHandler{}

	c.Subscribe(DUMMY_EVENT_1, &h)
	b.verifySubscribe(t, DUMMY_EVENT_1, c)
}

func TestPublish(t *testing.T) {
	b := &testBroker{}
	c := NewClient(b)

	e := dummyEvent{DUMMY_EVENT_1}

	c.Publish(e)
	b.verifyPublish(t, e)
}
