package curses

// Handles all output.

import (
	"log"

	"github.com/discoviking/roguemike/io"
	"github.com/rthornton128/goncurses"
)

var screen *goncurses.Window
var Input chan *io.UpdateBundle

func Init() error {
	s, err := goncurses.Init()
	screen = s
	if err != nil {
		return err
	}

	goncurses.Raw(true)
	goncurses.Echo(false)
	goncurses.Cursor(0)
	Input = make(chan *io.UpdateBundle, 1)

	log.Print("Starting output goroutine")
	go func() {
		for s := range Input {
			output(s)
		}
	}()

	return nil
}

func Term() {
	goncurses.End()
}

func output(u *io.UpdateBundle) {
	log.Print("Drawing update...")
	clearscreen()
	for _, e := range u.Entities {
		log.Printf("Drawing entity %#v", e)
		draw(e)
	}
	refresh()
}

func clearscreen() {
	screen.Erase()
}

func refresh() {
	screen.Refresh()
}

func draw(e *io.EntityData) {
	screen.MoveAddChar(e.Y, e.X, 'X')
}
