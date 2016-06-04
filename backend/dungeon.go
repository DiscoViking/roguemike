package backend

import (
	"log"
    "github.com/discoviking/roguemike/api"
)

type GameManager interface {
	Tick()
	GetState() (g *GameState)
	Data() *api.UpdateBundle
    SetInput (input <-chan Action)
}

type GameState struct {
	Entities []*Entity
	Actors   []*Actor
	Player   *Actor
}

func NewGameManager() GameManager {
	mgr := gameManager{}
	mgr.state = &GameState{}
	mgr.state.Player = NewPlayer()
	mgr.state.Entities = []*Entity{&mgr.state.Player.Entity}
	mgr.state.Actors = []*Actor{mgr.state.Player}
	return &mgr
}

type gameManager struct {
	state *GameState
}

func (mgr *gameManager) SetInput(input <-chan Action) {
    // TODO: Refactor this to avoid the type assertion.
    mgr.state.Player.Brain.(*InputBrain).SetInputChan(input)
}

func (mgr *gameManager) Tick() {
	for _, actor := range mgr.state.Actors {
		action := actor.ChooseAction(mgr.state)
		actor.Do(action)
	}
}

func (mgr *gameManager) GetState() (g *GameState) {
	return mgr.state
}

func (mgr *gameManager) Data() (bundle *api.UpdateBundle) {
	bundle = &api.UpdateBundle{}
	bundle.Entities = []*api.EntityData{}
	for _, entity := range mgr.state.Entities {
		log.Printf("Entity %#v", entity)
		bundle.Entities = append(bundle.Entities, entity.Data())
	}

	return bundle
}

