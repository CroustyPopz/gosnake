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

func (apple *Apple) resetPositions(mapSize int, snake *Snake) {
	if apple.eaten == true {
		apple.x = rand.Intn(mapSize - 1)
		apple.y = rand.Intn(mapSize - 1)
		for snake.isPresent(apple.x, apple.y) == true {
			apple.x = rand.Intn(mapSize - 1)
			apple.y = rand.Intn(mapSize - 1)
		}
		apple.eaten = false
	}
}

func (apple *Apple) Draw(snakeMap *SnakeMap, win *pixelgl.Window) {
	if apple.eaten == false {
		apple.mat = pixel.IM
		apple.mat = apple.mat.Moved(snakeMap.snakeMap[(snakeMap.mapSize*apple.x)+apple.y].Center())
		apple.sprite.Draw(win, apple.mat)
	}
}

func (apple *Apple) beEaten(snakeMap *SnakeMap, snake *Snake) {
	if apple.x == snake.sprites[0].x && apple.y == snake.sprites[0].y {
		apple.eaten = true
		snakeMap.score += 10
		snake.grow()
	}
}
