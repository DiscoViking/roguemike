package io

import "github.com/DiscoViking/roguemike/api"
import "github.com/DiscoViking/roguemike/backend"

type Manager struct {
	output chan<- *api.UpdateBundle
    playerInput chan backend.Action
}

func (mgr *Manager) Init() {
    mgr.playerInput = make(chan backend.Action, 1)
}

func (mgr *Manager) SetOutput(output chan<- *api.UpdateBundle) {
	mgr.output = output
}

func (mgr *Manager) GetPlayerInput() <-chan backend.Action {
    return mgr.playerInput
}

func (mgr *Manager) Update(bundle *api.UpdateBundle) {
	mgr.output <- bundle
}

func (mgr *Manager) HandleInput(clientAction api.ClientAction) {
    var action backend.Action
    switch clientAction.(type) {
    case api.MoveAction:
        clientMoveAction := clientAction.(api.MoveAction)
        action = &backend.ActionMove{DX:clientMoveAction.X, DY:clientMoveAction.Y}
    }

    mgr.playerInput <- action
}
