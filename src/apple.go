package main

import (
	"math/rand"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)

type Apple struct {
	eaten bool

	x int
	y int

	sprite *pixel.Sprite
	mat    pixel.Matrix
}

func NewApple(snake *Snake) *Apple {
	apple := Apple{eaten: true, x: 0, y: 0, sprite: snake.getFrame(0, 0), mat: pixel.IM}
	return &apple
}

func (apple *Apple) resetPositions() {
	if apple.eaten == true {
		apple.x = rand.Intn(9)
		apple.y = rand.Intn(9)
		apple.eaten = false
	}
}

func (apple *Apple) Draw(snakeMap *SnakeMap, win *pixelgl.Window) {
	apple.mat = pixel.IM
	apple.mat = apple.mat.Moved(snakeMap.snakeMap[(10*apple.x)+apple.y].Center())
	apple.sprite.Draw(win, apple.mat)
}
