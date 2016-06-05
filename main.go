package main

import (
	"log"
	"math/rand"
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

	for i := 0; i < 100; i++ {
		e := &backend.Entity{
			api.Coords{rand.Intn(50), rand.Intn(50)},
			0,
			api.TypeWall,
		}
		e.Init()
		game.Spawn(e)
	}

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
