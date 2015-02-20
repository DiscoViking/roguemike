package io

import "github.com/discoviking/roguemike/backend"

type UpdateBundle struct {
	Player   *PlayerData
	Entities []*EntityData
}

type PlayerData struct{}

type EntityData struct {
	Type backend.Type
	backend.Coord
}

func GetPlayerData(player *backend.Player) (data *PlayerData) {
	data = new(PlayerData)
	return data
}

func GetEntityData(entity *backend.Entity) (data *EntityData) {
	data = new(EntityData)
	data.X = entity.X
	data.Y = entity.Y
	data.Type = entity.Type

	return data
}

type ioManager struct {
	ioChan chan<- *UpdateBundle
}

func (mgr *ioManager) SetIOChan(ioChan chan<- *UpdateBundle) {
	mgr.ioChan = ioChan
}

func (mgr *ioManager) Update(g backend.GameManager) {
	bundle := UpdateBundle{}
	gameState := g.GetState()
	bundle.Player = GetPlayerData(gameState.Player)
	bundle.Entities = []*EntityData{}
	for _, entity := range gameState.Entities {
		bundle.Entities = append(bundle.Entities, GetEntityData(entity))
	}

	mgr.ioChan <- &bundle
}
