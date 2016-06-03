package api

type Coords struct {
	X int
	Y int
}

type EntityType uint32

const (
	TypeWall EntityType = iota
	TypePlayer
	TypeMonster
)

type UpdateBundle struct {
	Player   *PlayerData
	Entities []*EntityData
}

type PlayerData struct{}

type EntityData struct {
	Type EntityType
	Coords
}
