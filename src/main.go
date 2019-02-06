package main

import (
	"fmt"
	"image"
	_ "image/png"
	"os"
	"time"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/faiface/pixel/text"
	"golang.org/x/image/colornames"
	"golang.org/x/image/font/basicfont"
)

func DrawBG(bg *pixel.Sprite, win *pixelgl.Window) {
	mat := pixel.IM
	mat = mat.Moved(win.Bounds().Center())
	mat = mat.ScaledXY(win.Bounds().Center(), pixel.V(win.Bounds().H()/bg.Picture().Bounds().Size().X, win.Bounds().H()/bg.Picture().Bounds().Size().Y))
	bg.Draw(win, mat)
}

func initBG() *pixel.Sprite {
	pic, err := loadBG("assets/background.png")
	if err != nil {
		panic(err)
	}
	return pixel.NewSprite(pic, pic.Bounds())
}

func loadBG(path string) (pixel.Picture, error) {
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

func run() {

	cfg := pixelgl.WindowConfig{
		Title:  "Pixel Rocks!",
		Bounds: pixel.R(0, 0, 960, 960),
	}

	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}
	win.SetSmooth(true)

	bg := initBG()
	snake := NewSnake()
	apple := NewApple(snake)
	snake.initPositions(9, 9)
	snakeMap := NewSnakeMap(snake.frameSize, 15)
	atlas := text.NewAtlas(basicfont.Face7x13, text.ASCII)
	score := text.New(pixel.V(880, 900), atlas)
	score.Color = colornames.Black

	for !win.Closed() {
		snakeMap.dt = time.Since(snakeMap.last).Seconds()
		snake.initMatrix()

		// Detect move actions
		snakeMap.handleKeys(win, snake)

		// Turn management
		if snakeMap.dt > (0.5 - (float64(len(snake.sprites)) * 0.01)) {
			for snakeMap.gameover {
				// Detect restart control
				snakeMap.handleKeys(win, snake)

				mess := text.New(pixel.V(50, 500), atlas)
				mess.Color = colornames.Black
				fmt.Fprintln(mess, "PERDU!\nPresse Entrer pour recommencer")
				mess.Draw(win, pixel.IM.Scaled(mess.Orig, 4))
				win.Update()
			}

			apple.resetPositions(snakeMap.mapSize, snake)
			apple.beEaten(snakeMap, snake)
			score.Clear()
			fmt.Fprintln(score, snakeMap.score)
			snake.moveSnake(snakeMap) // change positions and head's sprites of the snake
			snake.setFrames()         // set the rights frames for the snake's body
			snakeMap.last = time.Now()
		}

		// win.Clear(colornames.Greenyellow)
		win.Clear(colornames.Firebrick)

		DrawBG(bg, win)
		apple.Draw(snakeMap, win)
		snake.Draw(snakeMap, win)
		score.Draw(win, pixel.IM.Scaled(score.Orig, 3))
		win.Update()
	}
}

func main() {
	pixelgl.Run(run)
}
