package backend

import "github.com/discoviking/roguemike/api"
import "github.com/discoviking/roguemike/events"

type InputBrain struct {
	input chan Action
}

func NewPlayer(eventManager *events.Manager) (player *Actor) {
	player = &Actor{}
	player.X = 10
	player.Y = 10
	player.Type = api.TypePlayer
	player.Brain = &InputBrain{make(chan Action, 1)}
    player.Brain.(*InputBrain).makeSubscriptions(eventManager)
	return player
}

func (b *InputBrain) ChooseAction(g *GameState) (action Action) {
    return <-b.input
}

func (b *InputBrain) makeSubscriptions(eventManager *events.Manager) {
    eventManager.Subscribe(
        events.Type("moveIntent"),
        events.Handler(func(e events.Event) {
            move := e.(api.MoveIntent)
            b.input <- &ActionMove{DX:move.X, DY:move.Y}
        }))
}
