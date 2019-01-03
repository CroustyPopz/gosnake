package main

import (
	"errors"
	"fmt"
	"image"
	"os"

	"github.com/faiface/pixel"
)

// This struct manage which frame to used and its matrix
type SnakePiece struct {
	sprite *pixel.Sprite
	mat    pixel.Matrix
}

// This struct is used to manage the all snake
type Snake struct {
	sprites   [1]SnakePiece // Rendered frames
	picture   pixel.Picture
	frames    []pixel.Rect
	frameSize float64
}

func NewSnake() *Snake {
	snake := Snake{frameSize: 64}
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
	for x := snake.picture.Bounds().Min.X; x < snake.picture.Bounds().Max.X; x += snake.frameSize {
		for y := snake.picture.Bounds().Min.Y; y < snake.picture.Bounds().Max.Y; y += snake.frameSize {
			snake.frames = append(snake.frames, pixel.R(x, y, x+snake.frameSize, y+snake.frameSize))
		}
	}
}

func (snake *Snake) getFrame(x int, y int) *pixel.Sprite {
	if x > 4 || x < 0 || y > 3 || y < 0 {
		panic(errors.New("Index not valid => out of range"))
	}

	fmt.Printf("%v\n", (4*x)+y)
	return pixel.NewSprite(snake.picture, snake.frames[(4*x)+y])
}
