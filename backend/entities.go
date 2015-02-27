package backend

import "sync"

var globalIdStore idStore

type Type uint8

type Entity struct {
	Coord
	ID   uint64
	Type Type
}

type Actor struct {
	Entity
	Brain
}

type Brain interface {
	ChooseAction(g *GameState) (action Action)
}

type idStore struct {
	sync.Mutex
	nextId uint64
}

func (entity *Entity) Init() {
	entity.ID = globalIdStore.NextID()
}

func (store *idStore) NextID() (id uint64) {
	store.Lock()
	defer store.Unlock()
	id = store.nextId
	store.nextId++
	return
}

func (actor *Actor) Do(action Action) {
	action.apply(actor)
}
