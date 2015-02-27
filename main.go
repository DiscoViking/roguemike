package main

import (
	"time"

	"github.com/discoviking/roguemike/backend"
	"github.com/discoviking/roguemike/io"
	"github.com/discoviking/roguemike/io/curses"
)

func main() {
	game := backend.NewGameManager()
	iomanager := &io.Manager{}
	curses.Init()
	defer curses.Term()
	iomanager.SetIOChan(curses.Input)

	t := time.NewTicker(1 * time.Second)
	count := 0

	for _ = range t.C {
		game.Tick()
		iomanager.Update(game)
		count++
		if count == 5 {
			break
		}
	}
}
