package main

import (
	"log"
	"os"

	"github.com/discoviking/roguemike/api"
	"github.com/discoviking/roguemike/backend"
	"github.com/discoviking/roguemike/events"
	"github.com/discoviking/roguemike/io/curses"
)

func main() {
	logFile, err := os.Create("roguemike.log")
	if err != nil {
		log.Fatal(err)
	}
	defer logFile.Close()
	log.SetOutput(logFile)

	eventsManager := events.NewManager()
	game := backend.NewGameManager(eventsManager)

	e := &backend.Entity{
		api.Coords{5, 5},
		0,
		api.TypeWall,
	}
	e.Init()
	game.Spawn(e)

	// Begin frontend loop.
	curses.Init(eventsManager)
	defer curses.Term()

	// Begin backend loop.
	go game.Loop()

	// Block until a 'quit' event is sent.
	quit := make(chan bool, 1)
	eventsManager.Subscribe(api.EventQuit, func(e events.Event) {
		quit <- true
	})
	<-quit
}
