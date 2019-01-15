package main

import (
	"errors"
	"image"
	"os"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
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
	snake.sprites[2] = SnakePiece{x: x, y: y - 2, sprite: snake.getFrame(3, 1), mat: pixel.IM}
}

func (snake *Snake) initMatrix() {
	for i := 0; i < len(snake.sprites); i++ {
		snake.sprites[i].mat = pixel.IM
	}
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

	return pixel.NewSprite(snake.picture, snake.frames[(4*x)+y])
}

func (snake *Snake) Draw(snakeMap *SnakeMap, win *pixelgl.Window) {
	for i := 0; i < len(snake.sprites); i++ {
		x := snake.sprites[i].x
		y := snake.sprites[i].y

		snake.sprites[i].mat = snake.sprites[i].mat.Moved(snakeMap.snakeMap[(10*x)+y].Center())
		snake.sprites[i].sprite.Draw(win, snake.sprites[i].mat)
	}
}

// get the move value and loop through the snake pieces
func (snake *Snake) moveSnake(snakeMap *SnakeMap) {
	snake.Move(1, snake.sprites[0].x, snake.sprites[0].y, snake.sprites[0].sprite)
	switch snakeMap.move {
	case -10:
		snake.sprites[0].sprite = snake.getFrame(3, 2)
		snake.sprites[0].x += -1
	case 10:
		snake.sprites[0].sprite = snake.getFrame(4, 3)
		snake.sprites[0].x += 1
	case 1:
		snake.sprites[0].sprite = snake.getFrame(3, 3)
		snake.sprites[0].y += 1
	case -1:
		snake.sprites[0].sprite = snake.getFrame(4, 2)
		snake.sprites[0].y += -1
	}
}

func (snake *Snake) Move(index int, prevX int, prevY int, prevSprite *pixel.Sprite) {
	if index == len(snake.sprites) {
		return
	}

	snake.Move(index+1, snake.sprites[index].x, snake.sprites[index].y, snake.sprites[index].sprite)
	snake.sprites[index].x = prevX
	snake.sprites[index].y = prevY

	// Tail
	if index == (len(snake.sprites) - 1) {
		// Up
		if snake.sprites[index-1].y > snake.sprites[index].y {
			snake.sprites[index].sprite = snake.getFrame(3, 1)

			// Down
		} else if snake.sprites[index-1].y < snake.sprites[index].y {
			snake.sprites[index].sprite = snake.getFrame(4, 0)

			// Right
		} else if snake.sprites[index-1].x > snake.sprites[index].x {
			snake.sprites[index].sprite = snake.getFrame(4, 1)

			// Left
		} else {
			snake.sprites[index].sprite = snake.getFrame(3, 0)
		}
		// Body
	} else {
		// Vertical Up
		if (snake.sprites[index-1].y > snake.sprites[index].y) && (snake.sprites[index].y > snake.sprites[index+1].y) {
			snake.sprites[index].sprite = snake.getFrame(2, 2)

			// Vertical Down
		} else if (snake.sprites[index-1].y < snake.sprites[index].y) && (snake.sprites[index].y < snake.sprites[index+1].y) {
			snake.sprites[index].sprite = snake.getFrame(2, 2)

			// Horizontal Right
		} else if (snake.sprites[index-1].x > snake.sprites[index].x) && (snake.sprites[index].x > snake.sprites[index+1].x) {
			snake.sprites[index].sprite = snake.getFrame(1, 3)

			// Horizontal Left
		} else if (snake.sprites[index-1].x < snake.sprites[index].x) && (snake.sprites[index].x < snake.sprites[index+1].x) {
			snake.sprites[index].sprite = snake.getFrame(1, 3)

			// Angle Up Right
		} else if (snake.sprites[index-1].x > snake.sprites[index].x) && (snake.sprites[index].y > snake.sprites[index+1].y) {
			snake.sprites[index].sprite = snake.getFrame(0, 3)

			// Angle Up Left
		} else if (snake.sprites[index-1].x < snake.sprites[index].x) && (snake.sprites[index].y > snake.sprites[index+1].y) {
			snake.sprites[index].sprite = snake.getFrame(2, 3)

			// Angle Down Right
		} else if (snake.sprites[index-1].x > snake.sprites[index].x) && (snake.sprites[index].y < snake.sprites[index+1].y) {
			snake.sprites[index].sprite = snake.getFrame(0, 2)

			// Angle Down Left
		} else {
			snake.sprites[index].sprite = snake.getFrame(2, 1)
		}
	}
}

func (snake *Snake) selectFrame(index int) {
	// directionX := snake.sprites[index-1].x - snake.sprites[index+1].x
	// directionY := snake.sprites[index-1].y - snake.sprites[index+1].y

	// if directionX ==  && directionY == 0 {

	// }
}
