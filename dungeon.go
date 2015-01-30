package backend

type IOManager interface {
	update(g GameState)
	setInputChan(c chan<- interface{})
}

type GameManager interface {
	tick(g GameState)
}

type GameState struct {
	entities []Entity
	player   Player
}

type Coord struct {
	X int
	Y int
}
