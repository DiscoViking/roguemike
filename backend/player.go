package backend

type Player struct {
	Actor
	input <-chan Action
}

func NewPlayer() (player *Player) {
	player = new(Player)
	player.Type = TypePlayer
	return player
}

func (player *Player) SetInputChan(input <-chan Action) {
	player.input = input
}

func (player *Player) ChooseAction(g *GameState) (action Action) {
	return <-player.input
}
