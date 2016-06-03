package io

import "github.com/discoviking/roguemike/api"

type Manager struct {
	output chan<- *api.UpdateBundle
}

func (mgr *Manager) SetOutput(output chan<- *api.UpdateBundle) {
	mgr.output = output
}

func (mgr *Manager) Update(bundle *api.UpdateBundle) {
	mgr.output <- bundle
}
