package backend

import (
	"log"
)

func transformOctant(row, col, oct int) (int, int) {
	switch oct {
	case 0:
		row, col = row, col
	case 1:
		row, col = col, row
	case 2:
		row, col = -col, row
	case 3:
		row, col = -row, col
	case 4:
		row, col = -row, -col
	case 5:
		row, col = -col, -row
	case 6:
		row, col = col, -row
	case 7:
		row, col = row, -col
	}

	return row, col
}

// Calculate which entities a given entity can see.
func getVisible(source *Entity, targets []*Entity, viewDistance int) []*Entity {
	visible := make([]*Entity, 0, len(targets))

	for oct := 0; oct < 8; oct++ {
		visible = append(visible, visibleOct(source, targets, viewDistance, oct)...)
	}

	return visible
}

func getTile(x, y int, targets []*Entity) *Entity {
	for _, e := range targets {
		if e.X == x && e.Y == y {
			return e
		}
	}

	return nil
}

type lineShadow struct {
	start float32
	end   float32
}

func (s lineShadow) occludes(row, col int) bool {
	log.Printf("Does %v occlude %v, %v?", s, row, col)
	if s.start*float32(row+1) <= float32(col) && s.end*float32(row+1) >= float32(col+1) {
		log.Print("yes")
		return true
	}
	log.Print("no")
	return false
}

func addShadow(row, col int, shadows []lineShadow) []lineShadow {
	this := lineShadow{
		start: float32(col) / float32(row+1),
		end:   float32(col+1) / float32(row+1),
	}
	log.Printf("Adding shadow: %v", this)

	// Assuming shadows are already ordered by start, and non-overlapping.
	// Find correct place to insert.
	for ix, s := range shadows {
		if s.start <= this.start {
			// Insert in order.
			log.Printf("Inserting at position %v", ix)
			shadows = append(shadows, this)
			copy(shadows[ix+2:], shadows[ix+1:])
			shadows[ix+1] = this
			break
		}
	}

	if len(shadows) == 0 {
		log.Printf("This is the first shadow")
		shadows = append(shadows, this)
	}

	// Merge any overlapping.
	newShadows := make([]lineShadow, 0, len(shadows))
	for _, s := range shadows {
		if len(newShadows) == 0 {
			newShadows = append(newShadows, s)
			continue
		}

		if s.start <= newShadows[len(newShadows)-1].end {
			newShadows[len(newShadows)-1].end = s.end
		} else {
			newShadows = append(newShadows, s)
		}
	}

	return newShadows
}

func visibleOct(source *Entity, targets []*Entity, viewDistance int, oct int) []*Entity {
	visible := make([]*Entity, 0, len(targets))
	X := source.X
	Y := source.Y

	shadows := make([]lineShadow, 0, viewDistance)

	for row := 1; row <= viewDistance; row++ {
		for col := 0; col <= row; col++ {
			//First, check if we can see this tile.
			canSee := true
			for _, s := range shadows {
				if s.occludes(row, col) {
					canSee = false
					break
				}
			}

			// If not, go to next tile.
			if !canSee {
				continue
			}

			// If so, check if there's an entity there.
			// If so, add it to the visible list and recalculate using its shadow.
			transRow, transCol := transformOctant(row, col, oct)
			if e := getTile(X+transCol, Y+transRow, targets); e != nil {
				visible = append(visible, e)
				shadows = addShadow(row, col, shadows)
				log.Printf("New shadows: %v", shadows)
			}
		}
	}

	return visible
}
