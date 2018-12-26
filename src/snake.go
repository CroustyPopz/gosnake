package main

import (
	"errors"
	"image"
	"os"

	"github.com/faiface/pixel"
)

type Snake struct {
	length     int
	spriteSize float64

	picture pixel.Picture
	frames  []pixel.Rect
}

func NewSnake() *Snake {
	snake := Snake{length: 3}
	snake.loadPicture()
	snake.setSnakeFrames()
	return &snake
}

func (snake *Snake) loadPicture() {
	file, err := os.Open("assets/snake.png")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	img, _, err := image.Decode(file)
	if err != nil {
		panic(err)
	}
	snake.picture = pixel.PictureDataFromImage(img)
}

func (snake *Snake) setSnakeFrames() {
	snake.spriteSize = snake.picture.Bounds().Size().X / 5

	for x := snake.picture.Bounds().Min.X; x < snake.picture.Bounds().Max.X; x += snake.spriteSize {
		for y := snake.picture.Bounds().Min.Y; y < snake.picture.Bounds().Max.Y; y += snake.spriteSize {
			snake.frames = append(snake.frames, pixel.R(x, y, x+snake.spriteSize, y+snake.spriteSize))
		}
	}
}

func (snake *Snake) getFrame(x int, y int) *pixel.Sprite {
	if x > 4 || x < 0 || y > 3 || y < 0 {
		panic(errors.New("Index not valid"))
	}

	sprite := pixel.NewSprite(snake.picture, snake.frames[(4*x)+y])
	return sprite
}
