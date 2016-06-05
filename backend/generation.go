package backend

// Dungeon Generator

import (
	"github.com/discoviking/roguemike/api"
	"math/rand"
)

// Make these not consts later.
const (
	maxRoomWidth  = 15
	maxRoomHeight = 10
	minRoomWidth  = 3
	minRoomHeight = 3
)

const (
	genEmpty = iota
	genWall
)

type room struct {
	x int
	y int
	w int
	h int
}

func GenerateNatural(width, height, openness int) []*Entity {
	grid := make([][]int, 0, height)
	for i := 0; i < height; i++ {
		row := make([]int, width)
		for j := 0; j < width; j++ {
			row[j] = genWall
		}
		grid = append(grid, row)
	}

	// Random walks to carve paths.
	for i := 0; i < openness; i++ {
		x := rand.Intn(width)
		y := rand.Intn(height)
		for j := 0; j < 100; j++ {
			x += rand.Intn(3) - 1
			y += rand.Intn(3) - 1
			if x >= width || y >= height || x < 0 || y < 0 {
				continue
			}
			grid[y][x] = genEmpty
		}
	}

	// Clean up unnecessary walls.
	unnecessary := [][]int{}
	for r := 0; r < height; r++ {
		for c := 0; c < width; c++ {
			if grid[r][c] != genWall {
				continue
			}

			// Check if any surrounding space is empty.
			necessary := false
			for dr := -1; dr <= 1; dr++ {
				for dc := -1; dc <= 1; dc++ {
					if r+dr < 0 || r+dr >= height || c+dc < 0 || c+dc >= width {
						continue
					}
					if grid[r+dr][c+dc] == genEmpty {
						necessary = true
						break
					}
				}
			}
			if !necessary {
				unnecessary = append(unnecessary, []int{r, c})
			}
		}
	}

	// Convert to entities.
	entities := []*Entity{}
	for i, row := range grid {
		for j, cell := range row {
			if cell == genWall {
				e := &Entity{
					api.Coords{i, j},
					0,
					api.TypeWall,
				}
				e.Init()
				entities = append(entities, e)
			}
		}
	}
	return entities
}

func Generate(width, height, openness int) []*Entity {
	// Firstly, generate some non-overlapping rooms.

	// Should calculate this somehow from width, height and openness.
	numRooms := 10

	rooms := make([]room, 0, numRooms)
	for len(rooms) < numRooms {
		room := randRoom(width, height, maxRoomWidth, maxRoomHeight)
		isGood := true
		for _, r := range rooms {
			if r.intersects(room) {
				isGood = false
				break
			}
		}

		if isGood {
			rooms = append(rooms, room)
		}
	}

	entities := make([]*Entity, 0, 1000)
	for _, r := range rooms {
		entities = append(entities, r.toWalls()...)
	}

	return entities
}

func randRoom(floorWidth, floorHeight, maxWidth, maxHeight int) room {
	w := rand.Intn(maxWidth) + minRoomWidth
	h := rand.Intn(maxHeight) + minRoomHeight
	return room{
		x: rand.Intn(floorWidth - w),
		y: rand.Intn(floorHeight - h),
		w: w,
		h: h,
	}
}

func (r room) intersects(other room) bool {
	// Are we to the right?
	if r.x > other.x+other.w {
		return false
	}

	// To the left?
	if r.x+r.w < other.x {
		return false
	}

	// Above?
	if r.y > other.y+other.h {
		return false
	}

	// Below?
	if r.y+r.h < other.y {
		return false
	}

	return true
}

func (r room) toWalls() []*Entity {
	walls := make([]*Entity, 0, r.w*2+r.h*2+4)

	// Top and bottom walls.
	for x := r.x; x < r.x+r.w; x++ {
		e := &Entity{
			api.Coords{x, r.y - 1},
			0,
			api.TypeWall,
		}
		e.Init()
		walls = append(walls, e)

		e = &Entity{
			api.Coords{x, r.y + r.h},
			0,
			api.TypeWall,
		}
		e.Init()
		walls = append(walls, e)
	}

	// Side walls.
	for y := r.y - 1; y < r.y+r.h+1; y++ {
		e := &Entity{
			api.Coords{r.x - 1, y},
			0,
			api.TypeWall,
		}
		e.Init()
		walls = append(walls, e)

		e = &Entity{
			api.Coords{r.x + r.w, y},
			0,
			api.TypeWall,
		}
		e.Init()
		walls = append(walls, e)
	}

	return walls
}
