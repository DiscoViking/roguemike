package backend

import "sync"

var globalIdStore idStore

type Entity struct {
	Position Coord
	ID       uint64
}

type Actor struct {
	Entity
	brain Brain
}

type Brain interface {
	// TODO
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
