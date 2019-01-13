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
	x int
	y int

	sprite *pixel.Sprite
	mat    pixel.Matrix
}

// This struct is used to manage the all snake
type Snake struct {
	sprites   [3]SnakePiece // Rendered frames
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

func (snake *Snake) initPositions(x int, y int) {
	snake.sprites[0] = SnakePiece{x: x, y: y, sprite: snake.getFrame(3, 3), mat: pixel.IM}
	snake.sprites[1] = SnakePiece{x: x, y: y - 1, sprite: snake.getFrame(2, 2), mat: pixel.IM}
	snake.sprites[2] = SnakePiece{x: 0, y: y - 2, sprite: snake.getFrame(3, 1), mat: pixel.IM}
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
