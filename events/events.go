package events

var global = NewManager()

type Type string

type Event interface {
	Type() Type
}

type Handler func(e Event)

type Manager struct {
	subscriptions map[Type][]Handler
}

func NewManager() *Manager {
	m := Manager{}
	m.subscriptions = make(map[Type][]Handler, 0)
	return &m
}

func (m *Manager) Subscribe(t Type, h Handler) {
	_, ok := m.subscriptions[t]
	if !ok {
		m.subscriptions[t] = make([]Handler, 0, 1)
	}

	m.subscriptions[t] = append(m.subscriptions[t], h)
}

func (m *Manager) Publish(e Event) {
	t := e.Type()
	if subscribers, ok := m.subscriptions[t]; ok {
		for _, s := range subscribers {
			s(e)
		}
	}
}

func Subscribe(t Type, h Handler) {
	global.Subscribe(t, h)
}

func Publish(e Event) {
	global.Publish(e)
}
