package main

import (
	_ "image/png"
	"os"
	"time"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
)

func run() {

	cfg := pixelgl.WindowConfig{
		Title:  "Pixel Rocks!",
		Bounds: pixel.R(0, 0, 640, 640),
	}

	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}
	win.SetSmooth(true)

	snake := NewSnake()
	apple := NewApple(snake)
	snake.initPositions(4, 4)
	snakeMap := NewSnakeMap(snake.frameSize, 10)

	for !win.Closed() {
		snakeMap.dt = time.Since(snakeMap.last).Seconds()
		snake.initMatrix()

		// Detect move actions
		snakeMap.handleKeys(win, snake)

		// Turn management
		if snakeMap.dt > 0.5 {
			if snakeMap.gameover {
				os.Exit(0)
			}
			apple.resetPositions()
			apple.beEaten(snake)
			snake.moveSnake(snakeMap) // change positions and head's sprites of the snake
			snake.setFrames()         // set the rights frames for the snake's body
			snakeMap.last = time.Now()
		}

		// win.Clear(colornames.Greenyellow)
		win.Clear(colornames.Firebrick)

		apple.Draw(snakeMap, win)
		snake.Draw(snakeMap, win)
		win.Update()
	}
}

func main() {
	pixelgl.Run(run)
}
