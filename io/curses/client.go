package curses

// Implementation of a RogueMike game client in curses.

import (
	"log"

	"github.com/discoviking/roguemike/api"
	"github.com/discoviking/roguemike/events"
	"github.com/rthornton128/goncurses"
)

var screen *goncurses.Window
var entityHistory map[api.Coords]api.EntityData

func Init(eventManager *events.Manager) error {
	s, err := goncurses.Init()
	screen = s
	if err != nil {
		return err
	}

	goncurses.Raw(true)
	goncurses.Echo(false)
	goncurses.Cursor(0)
	screen.Keypad(true)

	createEventSubscriptions(eventManager)
	go handleInput(eventManager)

	entityHistory = map[api.Coords]api.EntityData{}

	return nil
}

func Term() {
	goncurses.End()
}

func output(u *api.WorldUpdate) {
	log.Print("Drawing update...")
	clearscreen()
	for _, e := range entityHistory {
		log.Printf("Drawing historic entity %#v", &e)
		draw(&e, true)
	}
	for _, e := range u.Entities {
		log.Printf("Drawing entity %#v", &e)
		draw(&e, false)
		if e.Type == api.TypeWall {
			entityHistory[e.Coords] = e
		}
	}
	refresh()
}

func createEventSubscriptions(eventManager *events.Manager) {
	eventManager.Subscribe(
		api.EventWorldUpdate,
		events.Handler(func(e events.Event) {
			update := e.(api.WorldUpdate)
			output(&update)
		}))
}

func handleInput(eventManager *events.Manager) {
	for {
		c := screen.GetChar()
		switch c {
		case 'w', goncurses.KEY_UP:
			eventManager.Publish(api.MoveIntent{X: 0, Y: -1})
		case 'a', goncurses.KEY_LEFT:
			eventManager.Publish(api.MoveIntent{X: -1, Y: 0})
		case 'd', goncurses.KEY_RIGHT:
			eventManager.Publish(api.MoveIntent{X: 1, Y: 0})
		case 's', goncurses.KEY_DOWN:
			eventManager.Publish(api.MoveIntent{X: 0, Y: 1})
		case 'q':
			eventManager.Publish(api.Quit{})
		}
	}
}

func clearscreen() {
	screen.Erase()
}

func refresh() {
	screen.Refresh()
}

func draw(e *api.EntityData, history bool) {
	mod := goncurses.A_BOLD
	if history {
		mod = 0
	}
	switch e.Type {
	case api.TypeWall:
		screen.MoveAddChar(e.Y, e.X, goncurses.Char('X'|mod))
	case api.TypePlayer:
		screen.MoveAddChar(e.Y, e.X, '*')
	case api.TypeMonster:
		screen.MoveAddChar(e.Y, e.X, '@')
	}
}
