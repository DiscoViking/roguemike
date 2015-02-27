package backend

import "github.com/discoviking/roguemike/io"

type GameManager interface {
	Tick()
	GetState() (g *GameState)
	Data() *io.UpdateBundle
}

type GameState struct {
	Entities []*Entity
	Actors   []*Actor
	Player   *Player
}

func NewGameManager() GameManager {
	mgr := gameManager{}
	mgr.state = &GameState{}
	mgr.state.Player = NewPlayer()
	mgr.state.Entities = []*Entity{&mgr.state.Player.Entity}
	mgr.state.Actors = []*Actor{&mgr.state.Player.Actor}
	return &mgr
}

type gameManager struct {
	state *GameState
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

func (mgr *gameManager) Data() (bundle *io.UpdateBundle) {
	bundle = &io.UpdateBundle{}
	bundle.Player = mgr.state.Player.Data()
	bundle.Entities = []*io.EntityData{}
	for _, entity := range mgr.state.Entities {
		bundle.Entities = append(bundle.Entities, entity.Data())
	}

	return bundle
}
