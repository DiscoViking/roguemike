package events

type Type int

type Event interface {
	Type() Type
}

type Handler interface {
	Handle(e Event)
}

type Subscribable interface {
	Subscribe(t Type, h Handler)
}

type Publisher interface {
	Publish(e Event)
}

type Broker interface {
	Subscribable
	Publisher
}

type Client interface {
	Broker
	Handler
}

type HandlerFunc func(e Event)

func (f HandlerFunc) Handle(e Event) {
	f(e)
}
