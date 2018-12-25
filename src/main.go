package main

import (
	"errors"

	"image"
	"os"

	_ "image/png"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
)

func loadPicture(path string) (pixel.Picture, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	img, _, err := image.Decode(file)
	if err != nil {
		return nil, err
	}
	return pixel.PictureDataFromImage(img), nil
}

func getSnakeFrames(pic pixel.Picture) []pixel.Rect {
	var snakeFrames []pixel.Rect
	spriteLength := pic.Bounds().Size().X / 5

	for x := pic.Bounds().Min.X; x < pic.Bounds().Max.X; x += spriteLength {
		for y := pic.Bounds().Min.Y; y < pic.Bounds().Max.Y; y += spriteLength {
			snakeFrames = append(snakeFrames, pixel.R(x, y, x+spriteLength, y+spriteLength))
		}
	}

	return snakeFrames
}

func getSprite(pic pixel.Picture, snakeFrames []pixel.Rect, x int, y int) (*pixel.Sprite, error) {
	if x > 4 || x < 0 || y > 3 || y < 0 {
		return nil, errors.New("Index not valid")
	}

	sprite := pixel.NewSprite(pic, snakeFrames[(4*x)+y])

	return sprite, nil
}

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

	pic, err := loadPicture("assets/snake.png")
	snakeFrames := getSnakeFrames(pic)
	if err != nil {
		panic(err)
	}

	sprite, err := getSprite(pic, snakeFrames, 3, 3)

	for !win.Closed() {
		if win.JustPressed(pixelgl.KeyLeft) {
			sprite, err = getSprite(pic, snakeFrames, 3, 2)
		}
		if win.JustPressed(pixelgl.KeyRight) {
			sprite, err = getSprite(pic, snakeFrames, 4, 3)
		}
		if win.JustPressed(pixelgl.KeyUp) {
			sprite, err = getSprite(pic, snakeFrames, 3, 3)
		}
		if win.JustPressed(pixelgl.KeyDown) {
			sprite, err = getSprite(pic, snakeFrames, 4, 2)
		}

		// win.Clear(colornames.Greenyellow)
		win.Clear(colornames.Firebrick)

		mat := pixel.IM
		mat = mat.Moved(win.Bounds().Center())
		sprite.Draw(win, mat)

		win.Update()
	}
}

func main() {
	pixelgl.Run(run)
}
