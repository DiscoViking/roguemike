package main

import (
	"time"
)

func main() {
	err := InitOutput()
	if err != nil {
		panic(err)
	}
	defer TermOutput()

	e := Entity{
		Coord{15, 18},
		1,
	}

	Output([]Entity{e})

	time.Sleep(2 * time.Second)
}
