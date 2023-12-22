package game

import (
	"errors"
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"goSnake/pkg/config"
	"goSnake/pkg/image_manager"
	"goSnake/pkg/input"
	"goSnake/pkg/snake"
	"goSnake/pkg/utils/vector"
	"image/color"
	"math/rand"
	"time"
)

const tileSize = 32

type Game struct {
	speed int

	grid            *ebiten.Image
	queuedDirection int
	direction       int
	snake           *snake.Snake

	foodPos vector.Vector

	mainTicker *time.Ticker
}

func New() *Game {
	var g Game

	g.speed = 10
	g.grid = image_manager.Grid()
	g.snake = snake.New(vector.Vector{
		X: 32 * 5,
		Y: 32 * 5,
	})
	g.mainTicker = time.NewTicker(calcTick(g.speed))
	g.newFood()
	return &g
}

func (g *Game) Update() error {
	input.Update()

	switch g.direction {
	case input.Left, input.Right:
		if input.DirectionV() != input.No {
			g.queuedDirection = input.DirectionV()
		}
	case input.Up, input.Down:
		if input.DirectionH() != input.No {
			g.queuedDirection = input.DirectionH()
		}
	}

	select {
	case <-g.mainTicker.C:
		if g.queuedDirection != input.No {
			g.direction = g.queuedDirection
		}
		if g.direction == input.No {
			g.direction = input.Right
		}
		var newPos vector.Vector
		switch g.direction {
		case input.Left:
			newPos = vector.Vector{
				X: g.snake.Pos.X - 32,
				Y: g.snake.Pos.Y,
			}
		case input.Right:
			newPos = vector.Vector{
				X: g.snake.Pos.X + 32,
				Y: g.snake.Pos.Y,
			}
		case input.Up:
			newPos = vector.Vector{
				X: g.snake.Pos.X,
				Y: g.snake.Pos.Y - 32,
			}
		case input.Down:
			newPos = vector.Vector{
				X: g.snake.Pos.X,
				Y: g.snake.Pos.Y + 32,
			}
		}
		switch {
		case newPos.X < 0:
			newPos.X = config.FieldWidth - config.TileSize
		case newPos.Y < 0:
			newPos.Y = config.FieldHeight - config.TileSize
		case newPos.X+config.TileSize > config.FieldWidth:
			newPos.X = 0
		case newPos.Y+config.TileSize > config.FieldHeight:
			newPos.Y = 0
		}

		s := g.snake.Next
		for s != nil {
			if s.Pos == newPos {
				return errors.New("you died")
			}
			s = s.Next
		}

		g.snake.Move(newPos)
		if g.snake.Pos == g.foodPos {
			g.snake.Grow()
			g.newFood()
		}
	default:
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	snakeField := ebiten.NewImage(config.FieldWidth+1, config.FieldHeight+1)
	snakeField.Fill(color.RGBA{10, 10, 10, 0})

	snakeField.DrawImage(g.grid, nil)
	if g.snake != nil {
		g.snake.Draw(snakeField, calcTick(g.speed))
	}

	var foodDrawOptions ebiten.DrawImageOptions
	foodDrawOptions.GeoM.Translate(g.foodPos.X, g.foodPos.Y)
	snakeField.DrawImage(image_manager.Food(), &foodDrawOptions)

	var snakeFieldDrawingOptions ebiten.DrawImageOptions
	snakeFieldDrawingOptions.GeoM.Translate(config.FieldLeft, config.FieldTop)
	screen.DrawImage(snakeField, &snakeFieldDrawingOptions)

	ebitenutil.DebugPrint(screen, fmt.Sprintf("FPS: %d", int(ebiten.ActualFPS())))
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("TPS: %d", int(ebiten.ActualTPS())), 0, 10)
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("speed: %d", g.speed), 0, 20)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideHeight
}

func (g *Game) newFood() {
	newFoodPos := vector.Vector{
		X: float64(rand.Intn(config.FieldWidth/config.TileSize) * config.TileSize),
		Y: float64(rand.Intn(config.FieldHeight/config.TileSize) * config.TileSize),
	}

	s := g.snake
	for s != nil {
		if s.Pos == newFoodPos {
			fmt.Println("collision")
			newFoodPos = vector.Vector{
				X: float64(rand.Intn(config.FieldWidth/config.TileSize) * config.TileSize),
				Y: float64(rand.Intn(config.FieldHeight/config.TileSize) * config.TileSize),
			}
		}
		s = s.Next
	}
	g.foodPos = newFoodPos
}

func (g *Game) SpeedUp() {
	g.speed++
	g.mainTicker = time.NewTicker(calcTick(g.speed))
}

func calcTick(speed int) time.Duration {
	return (time.Second) / time.Duration(speed)
}
