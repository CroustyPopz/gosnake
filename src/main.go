package main

import (
	_ "image/png"
	"time"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
)

func buildSnakeMap(tileSize float64, mapSize int) []pixel.Rect {
	var snakeMap []pixel.Rect

	for x := 0; x < mapSize; x++ {
		for y := 0; y < mapSize; y++ {
			r := pixel.R(float64(x)*tileSize, float64(y)*tileSize, (float64(x)*tileSize)+tileSize, (float64(y)*tileSize)+tileSize)
			snakeMap = append(snakeMap, r)
		}
	}

	return snakeMap
}

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
	sprite := snake.getFrame(3, 3)
	snakeMap := buildSnakeMap(snake.frameSize, 10)

	index := 45
	move := 1
	last := time.Now()

	for !win.Closed() {
		dt := time.Since(last).Seconds()
		mat := pixel.IM

		if win.JustPressed(pixelgl.KeyLeft) {
			sprite = snake.getFrame(3, 2)
			move = -10
		}
		if win.JustPressed(pixelgl.KeyRight) {
			sprite = snake.getFrame(4, 3)
			move = 10
		}
		if win.JustPressed(pixelgl.KeyUp) {
			sprite = snake.getFrame(3, 3)
			move = 1
		}
		if win.JustPressed(pixelgl.KeyDown) {
			sprite = snake.getFrame(4, 2)
			move = -1
		}

		// win.Clear(colornames.Greenyellow)
		if dt > 0.5 {
			index += move
			last = time.Now()
		}

		mat = mat.Moved(snakeMap[index].Center())
		win.Clear(colornames.Firebrick)
		sprite.Draw(win, mat)
		win.Update()
	}
}

func main() {
	pixelgl.Run(run)
}
