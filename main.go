package main

import (
	"log"
	"os"

	"github.com/discoviking/roguemike/api"
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
	iomanager.Init()
	curses.Init()
	defer curses.Term()

	iomanager.SetOutput(curses.Input)
	game.SetInput(iomanager.GetPlayerInput())

	iomanager.Update(game.Data())
	for action := range curses.Output {
        _, shouldQuit := (*action).(api.QuitAction)
        if (shouldQuit) {
            break
        }

		iomanager.HandleInput(*action)
		log.Print("Updating IO")
		game.Tick()
		iomanager.Update(game.Data())
	}
}
