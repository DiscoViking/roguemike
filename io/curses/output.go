package curses

// Handles all output.

import (
	"code.google.com/p/goncurses"
	"github.com/discoviking/roguemike"
	"github.com/discoviking/roguemike/io"
)

var screen *goncurses.Window
var Input chan *io.UpdateBundle

func New() error {
	s, err := goncurses.Init()
	screen = s
	if err != nil {
		return err
	}

	goncurses.Raw(true)
	goncurses.Echo(false)
	goncurses.Cursor(0)

	go func() {
		for s := range Input {
			Output(s)
		}
	}()

	return nil
}

func Term() {
	goncurses.End()
}

func output(u *io.UpdateBundle) {
	clearscreen()
	for _, e := range u.Entities {
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

func draw(e EntityData) {
	screen.MoveAddChar(e.Y, e.X, 'X')
}
