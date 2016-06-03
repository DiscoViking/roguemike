package backend

import (
	"sync"

	"github.com/discoviking/roguemike/api"
)

var globalIdStore idStore

type Entity struct {
	api.Coords
	ID   uint64
	Type api.EntityType
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

func (entity *Entity) Data() *api.EntityData {
	data := &api.EntityData{}
	data.X = entity.X
	data.Y = entity.Y
	data.Type = entity.Type

	return data
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
