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

type Manager struct {
	output chan<- *UpdateBundle
	input  chan<- backend.Action
}

func (mgr *Manager) SetOutput(output chan<- *UpdateBundle) {
	mgr.output = output
}

func (mgr *Manager) SetInput(input chan<- backend.Action) {
	mgr.input = input
}

func (mgr *Manager) SendAction(action backend.Action) {
	mgr.input <- action
}

func (mgr *Manager) Update(g backend.GameManager) {
	bundle := UpdateBundle{}
	gameState := g.GetState()
	bundle.Player = GetPlayerData(gameState.Player)
	bundle.Entities = []*EntityData{}
	for _, entity := range gameState.Entities {
		bundle.Entities = append(bundle.Entities, GetEntityData(entity))
	}

	mgr.output <- &bundle
}
