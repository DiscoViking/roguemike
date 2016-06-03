package main

import (
	"log"
	"os"
	"time"

	"github.com/discoviking/roguemike/backend"
	"github.com/discoviking/roguemike/io"
	"github.com/discoviking/roguemike/io/curses"
)

func main() {
	logFile, err := os.Create("roguemike.log")
	if err != nil {
		log.Fatal(err)
	}
	defer logFile.Close()
	log.SetOutput(logFile)

	game := backend.NewGameManager()
	iomanager := &io.Manager{}
	curses.Init()
	defer curses.Term()
	iomanager.SetOutput(curses.Input)

	t := time.NewTicker(1 * time.Second)
	count := 0

	for _ = range t.C {
		log.Print("Updating IO")
		//game.Tick()
		iomanager.Update(game.Data())
		count++
		if count == 5 {
			break
		}
	}
}
