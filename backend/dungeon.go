package backend

type IOManager interface {
	Update(g *GameState)
	SetInputChan(c chan<- interface{})
}

type GameManager interface {
	Tick(g *GameState)
}

type GameState struct {
	Entities []Entity
	Actors   []Actor
	Player   Player
}

type Coord struct {
	X int
	Y int
}

type gameManager struct {
}

func (mgr *gameManager) Tick(g *GameState) {
	for _, actor := range g.Actors {
		action := actor.ChooseAction(g)
		actor.Do(action)
	}
}
