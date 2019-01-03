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
	snake.sprites[0].sprite = snake.getFrame(3, 3)
	snakeMap := NewSnakeMap(snake.frameSize, 10)

	for !win.Closed() {
		snakeMap.dt = time.Since(snakeMap.last).Seconds()
		mat := pixel.IM

		snakeMap.handleKeys(snake, win)

		// win.Clear(colornames.Greenyellow)
		if snakeMap.dt > 0.5 {
			snakeMap.index += snakeMap.move
			snakeMap.last = time.Now()
		}

		mat = mat.Moved(snakeMap.snakeMap[snakeMap.index].Center())
		win.Clear(colornames.Firebrick)
		snake.sprites[0].sprite.Draw(win, mat)
		win.Update()
	}
}

func main() {
	pixelgl.Run(run)
}
