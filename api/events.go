package api

import "github.com/DiscoViking/roguemike/events"

// Event Types
const (
	EventWorldUpdate events.Type = iota
	EventMoveIntent
	EventQuit
)

type WorldUpdate struct {
	Player   PlayerData
	Entities []EntityData
}

func (u WorldUpdate) Type() events.Type {
	return EventWorldUpdate
}

type MoveIntent struct {
	X int
	Y int
}

func (i MoveIntent) Type() events.Type {
	return EventMoveIntent
}

type Quit struct{}

func (i Quit) Type() events.Type {
	return EventQuit
}
