package backend

import "github.com/DiscoViking/roguemike/api"
import "github.com/DiscoViking/roguemike/events"

type InputBrain struct {
	input chan Action
}

func NewPlayer(broker events.Broker) (player *Actor) {
	player = &Actor{}
	player.X = 10
	player.Y = 10
	player.Type = api.TypePlayer
	player.Brain = &InputBrain{make(chan Action, 1)}
	player.Brain.(*InputBrain).makeSubscriptions(broker)
	return player
}

func (b *InputBrain) ChooseAction(g *GameState) (action Action) {
	return <-b.input
}

func (b *InputBrain) makeSubscriptions(broker events.Broker) {
	broker.Subscribe(
		events.Type(api.EventMoveIntent),
		events.HandlerFunc(func(e events.Event) {
			move := e.(api.MoveIntent)
			b.input <- &ActionMove{DX: move.X, DY: move.Y}
		}))
}
