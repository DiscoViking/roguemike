package backend

import "github.com/discoviking/roguemike/api"

type Action interface {
	apply(actor *Actor, state *GameState)
}

// Move the actor DX pixels right, and DY pixels down.
type ActionMove struct {
	DX int
	DY int
}

func (action *ActionMove) apply(actor *Actor, state *GameState) {
    if (state.IsTraversable(api.Coords{actor.X + action.DX, actor.Y + action.DY})) {
        actor.X += action.DX
        actor.Y += action.DY
    }
}
