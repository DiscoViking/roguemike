package main

// Handles all output.

import (
	"code.google.com/p/goncurses"
)

var screen *goncurses.Window

func InitOutput() error {
	s, err := goncurses.Init()
	screen = s
	if err != nil {
		return err
	}

	goncurses.Raw(true)
	goncurses.Echo(false)
	goncurses.Cursor(0)

	return nil
}

func TermOutput() {
	goncurses.End()
}

func Output(entities []Entity) {
	clearscreen()
	for _, e := range entities {
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

func draw(e Entity) {
	screen.MoveAddChar(e.Position.Y, e.Position.X, 'X')
}
