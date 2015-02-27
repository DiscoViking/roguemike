package backend

type GameManager interface {
	Tick(g *GameState)
	GetState() (g *GameState)
}

type GameState struct {
	Entities []*Entity
	Actors   []Actor
	Player   *Player
}

type Coord struct {
	X int
	Y int
}

type gameManager struct {
	ioChan <-chan *GameState
}

func (mgr *gameManager) Tick(g *GameState) {
	for _, actor := range g.Actors {
		action := actor.ChooseAction(g)
		actor.Do(action)
	}
}

func (mgr *gameManager) GetState() (g *GameState) {
	return g
}
