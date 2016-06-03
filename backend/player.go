package backend

import "github.com/discoviking/roguemike/api"

type InputBrain struct {
	input <-chan Action
}

func NewPlayer() (player *Actor) {
	player = &Actor{}
	player.X = 10
	player.Y = 10
	player.Type = api.TypePlayer
	player.Brain = &InputBrain{}
	return player
}

func (b *InputBrain) SetInputChan(input <-chan Action) {
	b.input = input
}

func (b *InputBrain) ChooseAction(g *GameState) (action Action) {
	return <-b.input
}
