package curses

// Implementation of a RogueMike game client in curses.

import (
	"log"

	"github.com/discoviking/roguemike/api"
	"github.com/rthornton128/goncurses"
)

var screen *goncurses.Window
var Input chan *api.UpdateBundle
var Output chan *api.ClientAction

func Init() error {
	s, err := goncurses.Init()
	screen = s
	if err != nil {
		return err
	}

	goncurses.Raw(true)
	goncurses.Echo(false)
	goncurses.Cursor(0)
    screen.Keypad(true)

	Input = make(chan *api.UpdateBundle, 1)
	Output = make(chan *api.ClientAction, 1)

	log.Print("Starting output goroutine")
	go func() {
		for s := range Input {
			output(s)
		}
	}()

    go handleInput()

	return nil
}

func Term() {
	goncurses.End()
}

func output(u *api.UpdateBundle) {
	log.Print("Drawing update...")
	clearscreen()
	for _, e := range u.Entities {
		log.Printf("Drawing entity %#v", e)
		draw(e)
	}
	refresh()
}

func handleInput() {
    for {
        var action api.ClientAction = nil;
        c := screen.GetChar()
        switch c {
        case 'w', goncurses.KEY_UP:
            action = api.MoveAction{X:0, Y:-1}
        case 'a', goncurses.KEY_LEFT:
            action = api.MoveAction{X:-1, Y:0}
        case 'd', goncurses.KEY_RIGHT:
            action = api.MoveAction{X:1, Y:0}
        case 's', goncurses.KEY_DOWN:
            action = api.MoveAction{X:0, Y:1}
        case 'q':
            action = api.QuitAction{}
        }

        if (action != nil) {
            Output <- &action
        }
    }
}

func clearscreen() {
	screen.Erase()
}

func refresh() {
	screen.Refresh()
}

func draw(e *api.EntityData) {
    switch e.Type {
    case api.TypeWall:
        screen.MoveAddChar(e.Y, e.X, 'X')
    case api.TypePlayer:
        screen.MoveAddChar(e.Y, e.X, '*')
    case api.TypeMonster:
        screen.MoveAddChar(e.Y, e.X, '@')
    }
}
