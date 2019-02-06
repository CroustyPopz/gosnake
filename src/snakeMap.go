package main

import (
	"fmt"
	"os"
	"time"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)

// This struct is used to manage the core game
type SnakeMap struct {
	gameover bool
	score    int

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
	snakeMap := SnakeMap{gameover: false, score: 0, mapSize: mapSize, move: 1, last: time.Now(), dt: 0}
	snakeMap.buildSnakeMap(tileSize)

	return &snakeMap
}

func (snakeMap *SnakeMap) handleKeys(win *pixelgl.Window, snake *Snake) {
	if win.Pressed(pixelgl.KeyEnter) && snakeMap.gameover == true {
		snakeMap.gameover = false
		snakeMap.score = 0
		snake.initPositions(9, 9)
		snakeMap.move = 1
	}
	if win.Pressed(pixelgl.KeyEscape) {
		os.Exit(0)
	}
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

func (snakeMap *SnakeMap) buildSnakeMap(tileSize float64) {
	for x := 0; x < snakeMap.mapSize; x++ {
		for y := 0; y < snakeMap.mapSize; y++ {
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
	return snakeMap.gameover
}
