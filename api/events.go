package api

import "github.com/discoviking/roguemike/events"

type WorldUpdate struct {
	Player   PlayerData
	Entities []EntityData
}

func (u WorldUpdate) Type() events.Type {
    return events.Type("worldupdate")
}

type MoveIntent struct {
    X int
    Y int
}

func (i MoveIntent) Type() events.Type {
    return events.Type("moveintent")
}

type Quit struct {}

func (i Quit) Type() events.Type {
    return events.Type("quit")
}
