package api

import "github.com/DiscoViking/roguemike/events"

// Event Types
const (
	EventWorldUpdate events.Type = iota
	EventMoveIntent
	EventQuit
	EventTick
)

const (
	Quit = basicEvent(EventQuit)
	Tick = basicEvent(EventTick)
)

type basicEvent events.Type

func (b basicEvent) Type() events.Type { return events.Type(b) }

type WorldUpdate struct {
	Player   PlayerData
	Entities []EntityData
}

func (u WorldUpdate) Type() events.Type { return EventWorldUpdate }

type MoveIntent struct {
	X int
	Y int
}

func (i MoveIntent) Type() events.Type { return EventMoveIntent }
