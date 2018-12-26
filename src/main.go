package main

import (
	_ "image/png"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
)

func run() {

	cfg := pixelgl.WindowConfig{
		Title:  "Pixel Rocks!",
		Bounds: pixel.R(0, 0, 1024, 768),
	}

	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}
	win.SetSmooth(true)

	snake := NewSnake()
	sprite := snake.getFrame(3, 3)

	mat := pixel.IM
	mat = mat.Moved(win.Bounds().Center())

	for !win.Closed() {
		if win.JustPressed(pixelgl.KeyLeft) {
			sprite = snake.getFrame(3, 2)
		}
		if win.JustPressed(pixelgl.KeyRight) {
			sprite = snake.getFrame(4, 3)
		}
		if win.JustPressed(pixelgl.KeyUp) {
			sprite = snake.getFrame(3, 3)
		}
		if win.JustPressed(pixelgl.KeyDown) {
			sprite = snake.getFrame(4, 2)
		}

		// win.Clear(colornames.Greenyellow)
		win.Clear(colornames.Firebrick)
		sprite.Draw(win, mat)
		win.Update()
	}
}

func main() {
	pixelgl.Run(run)
}
