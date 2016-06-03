package common

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
