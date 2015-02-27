package backend

type GameManager interface {
	Tick()
	GetState() (g *GameState)
}

type GameState struct {
	Entities []*Entity
	Actors   []*Actor
	Player   *Player
}

type Coord struct {
	X int
	Y int
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
