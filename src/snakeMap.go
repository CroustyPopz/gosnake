package main

import (
	"fmt"
	"time"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)

// This struct is used to manage the core game
type SnakeMap struct {
	gameover bool

	// MAP
	mapSize  int
	snakeMap []pixel.Rect

	// SNAKE POSITION
	move int

	// LOOP MANAGERS
	last time.Time
	dt   float64
}

func NewSnakeMap(tileSize float64, mapSize int) *SnakeMap {
	snakeMap := SnakeMap{gameover: false, mapSize: 10, move: 1, last: time.Now(), dt: 0}
	snakeMap.buildSnakeMap(tileSize, mapSize)

	return &snakeMap
}

func (snakeMap *SnakeMap) handleKeys(win *pixelgl.Window, snake *Snake) {
	if win.JustPressed(pixelgl.KeyLeft) {
		if snake.sprites[1].x == snake.sprites[0].x-1 && snake.sprites[1].y == snake.sprites[0].y {
			return
		}
		snakeMap.move = -10
	}
	if win.JustPressed(pixelgl.KeyRight) {
		if snake.sprites[1].x == snake.sprites[0].x+1 && snake.sprites[1].y == snake.sprites[0].y {
			return
		}
		snakeMap.move = 10
	}
	if win.JustPressed(pixelgl.KeyUp) {
		if snake.sprites[1].x == snake.sprites[0].x && snake.sprites[1].y == snake.sprites[0].y+1 {
			return
		}
		snakeMap.move = 1
	}
	if win.JustPressed(pixelgl.KeyDown) {
		if snake.sprites[1].x == snake.sprites[0].x && snake.sprites[1].y == snake.sprites[0].y-1 {
			return
		}
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

func (snakeMap *SnakeMap) isOutside(x int, y int) bool {
	if x < 0 || y < 0 || x >= snakeMap.mapSize || y >= snakeMap.mapSize {
		fmt.Print("You're out of the map! Game oveR...")
		return true
	}
	return false
}
