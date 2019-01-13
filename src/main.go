package main

import (
	_ "image/png"
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
	snake.initPositions(4, 4)
	snakeMap := NewSnakeMap(snake.frameSize, 10)

	for !win.Closed() {
		snakeMap.dt = time.Since(snakeMap.last).Seconds()
		snake.sprites[0].mat = pixel.IM

		snakeMap.handleKeys(snake, win)

		// win.Clear(colornames.Greenyellow)
		if snakeMap.dt > 0.5 {
			snakeMap.index += snakeMap.move
			snakeMap.moveSnake(snake)
			snakeMap.last = time.Now()
		}

		snake.sprites[0].mat = snake.sprites[0].mat.Moved(snakeMap.snakeMap[snakeMap.index].Center())
		win.Clear(colornames.Firebrick)
		snake.sprites[0].sprite.Draw(win, snake.sprites[0].mat)
		win.Update()
	}
}

func main() {
	pixelgl.Run(run)
}
