package backend

import (
	"log"
    "github.com/discoviking/roguemike/api"
    "github.com/discoviking/roguemike/events"
)

type GameManager interface {
	Tick()
    Loop()
	GetState() (g *GameState)
}

type GameState struct {
	Entities []*Entity
	Actors   []*Actor
	Player   *Actor
}

func NewGameManager(eventsManager *events.Manager) GameManager {
	mgr := gameManager{}
	mgr.state = &GameState{}
	mgr.state.Player = NewPlayer(eventsManager)
	mgr.state.Entities = []*Entity{&mgr.state.Player.Entity}
	mgr.state.Actors = []*Actor{mgr.state.Player}
    mgr.eventsManager = eventsManager
	return &mgr
}

type gameManager struct {
	state *GameState
    eventsManager *events.Manager
}

func (mgr *gameManager) Tick() {
	for _, actor := range mgr.state.Actors {
		action := actor.ChooseAction(mgr.state)
		action.apply(actor, mgr.state)
	}

    mgr.pushUpdate()
}

func (mgr *gameManager) Loop() {
    for {
        mgr.Tick()
    }
}

func (mgr *gameManager) GetState() (g *GameState) {
	return mgr.state
}

func (mgr *gameManager) pushUpdate() {
    update := api.WorldUpdate{}
	update.Entities = []api.EntityData{}
	for _, entity := range mgr.state.Entities {
		log.Printf("Entity %#v", entity)
        update.Entities = append(update.Entities, *entity.Data())
	}

    mgr.eventsManager.Publish(update)
}

func (state *GameState) IsTraversable(position api.Coords) bool {
    for _, entity := range state.Entities {
        if (entity.X == position.X && entity.Y == position.Y) {
            return false
        }
    }

    return true
}
