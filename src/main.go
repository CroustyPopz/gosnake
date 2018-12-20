package main

import (
	"errors"
	"fmt"

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

func getSprite(pic pixel.Picture, x float64, y float64) (*pixel.Sprite, error) {
	if x > 4 || x < 0 || y > 3 || y < 0 {
		return nil, errors.New("Index not valid")
	}
	spriteLength := pic.Bounds().Size().Len() / 5

	rect := pixel.R(x*spriteLength, y*spriteLength, (x+1)*spriteLength, (y+1)*spriteLength)
	sprite := pixel.NewSprite(pic, rect)
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

	pic, err := loadPicture("assets/snake.png")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%v", pic.Bounds())

	sprite, err := getSprite(pic, 0, 0)

	win.Clear(colornames.Greenyellow)
	sprite.Draw(win, pixel.IM.Moved(win.Bounds().Center()))
	for !win.Closed() {
		win.Update()
	}
}

func main() {
	pixelgl.Run(run)
}
