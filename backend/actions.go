package backend

type Action interface {
	apply(actor *Actor)
}

// Move the actor DX pixels right, and DY pixels down.
type ActionMove struct {
	DX int
	DY int
}

func (action *ActionMove) apply(actor *Actor) {
	actor.X += action.DX
	actor.Y += action.DY
}
