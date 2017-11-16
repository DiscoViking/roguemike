package main

import (
	"log"
	"os"

	"github.com/DiscoViking/roguemike/api"
	"github.com/DiscoViking/roguemike/backend"
	"github.com/DiscoViking/roguemike/events"
	"github.com/DiscoViking/roguemike/io/curses"
)

func main() {
	logFile, err := os.Create("roguemike.log")
	if err != nil {
		log.Fatal(err)
	}
	defer logFile.Close()
	log.SetOutput(logFile)

	broker := events.NewBroker()
	game := backend.NewGameManager(broker)

	e := &backend.Entity{
		api.Coords{5, 5},
		0,
		api.TypeWall,
	}
	e.Init()
	game.Spawn(e)

	// Begin frontend loop.
	curses.Init(broker)
	defer curses.Term()

	// Begin backend loop.
	go game.Loop()

	// Block until a 'quit' event is sent.
	quit := make(chan bool, 1)
	broker.Subscribe(api.EventQuit, events.HandlerFunc(func(e events.Event) {
		quit <- true
	}))
	<-quit
}
