package backend

import (
	"sync"

	"github.com/discoviking/roguemike/common"
	"github.com/discoviking/roguemike/io"
)

var globalIdStore idStore

type Entity struct {
	common.Coords
	ID   uint64
	Type io.EntityType
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

func (entity *Entity) Data() *io.EntityData {
	data := &io.EntityData{}
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
