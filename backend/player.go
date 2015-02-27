package backend

import "github.com/discoviking/roguemike/io"

type Player struct {
	Actor
	input <-chan Action
}

func NewPlayer() (player *Player) {
	player = new(Player)
	player.Type = TypePlayer
	return player
}

func (player *Player) Data() *io.PlayerData {
	return &io.PlayerData{}
}

func (player *Player) SetInputChan(input <-chan Action) {
	player.input = input
}

func (player *Player) ChooseAction(g *GameState) (action Action) {
	return <-player.input
}
