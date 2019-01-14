package main

import (
	"time"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)

// This struct is used to manage the core game
type SnakeMap struct {
	// MAP
	snakeMap []pixel.Rect

	// SNAKE POSITION
	index int
	move  int

	// LOOP MANAGERS
	last time.Time
	dt   float64
}

func NewSnakeMap(tileSize float64, mapSize int) *SnakeMap {
	snakeMap := SnakeMap{index: (mapSize * 4) + 4, move: 1, last: time.Now(), dt: 0}
	snakeMap.buildSnakeMap(tileSize, mapSize)

	return &snakeMap
}

func (snakeMap *SnakeMap) handleKeys(win *pixelgl.Window) {
	if win.JustPressed(pixelgl.KeyLeft) {
		snakeMap.move = -10
	}
	if win.JustPressed(pixelgl.KeyRight) {
		snakeMap.move = 10
	}
	if win.JustPressed(pixelgl.KeyUp) {
		snakeMap.move = 1
	}
	if win.JustPressed(pixelgl.KeyDown) {
		snakeMap.move = -1
	}
}

func (snakeMap *SnakeMap) buildSnakeMap(tileSize float64, mapSize int) {

	for x := 0; x < mapSize; x++ {
		for y := 0; y < mapSize; y++ {
			r := pixel.R(float64(x)*tileSize, float64(y)*tileSize, (float64(x)*tileSize)+tileSize, (float64(y)*tileSize)+tileSize)
			snakeMap.snakeMap = append(snakeMap.snakeMap, r)
		}
	}
}
