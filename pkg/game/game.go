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
	"goSnake/pkg/wall"
	"math/rand"
	"time"
)

const tileSize = 32

type Game struct {
	direction int
	snake     *snake.Snake
	walls     []*wall.Wall

	foodPos vector.Vector

	mainTicker *time.Ticker
}

func New() *Game {
	var g Game
	g.snake = snake.New(vector.Vector{
		X: 32,
		Y: 32,
	})
	g.walls = []*wall.Wall{
		wall.New(vector.Vector{
			X: 10 * config.TileSize,
			Y: 10 * config.TileSize,
		}),
		wall.New(vector.Vector{
			X: 15 * config.TileSize,
			Y: 4 * config.TileSize,
		}),
	}
	g.mainTicker = time.NewTicker(config.Tick)
	g.newFood()
	return &g
}

func (g *Game) Update() error {
	input.Update()

	select {
	case <-g.mainTicker.C:
		switch g.direction {
		case input.Left, input.Right:
			if input.DirectionV() != input.No {
				g.direction = input.DirectionV()
			}
		case input.Up, input.Down:
			if input.DirectionH() != input.No {
				g.direction = input.DirectionH()
			}
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
		if newPos.X < 0 ||
			newPos.Y < 0 ||
			newPos.X+config.TileSize > config.ScreenWidth ||
			newPos.Y+config.TileSize > config.ScreenHeight {
			return errors.New("you died")
		}

		for _, wall := range g.walls {
			if wall.IsCollides(g.snake.Pos) {
				return errors.New("you died")
			}
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
	if g.snake != nil {
		g.snake.Draw(screen)
	}

	var foodDrawOptions ebiten.DrawImageOptions
	foodDrawOptions.GeoM.Translate(g.foodPos.X, g.foodPos.Y)
	screen.DrawImage(image_manager.Food(), &foodDrawOptions)

	for _, wall := range g.walls {
		wall.Draw(screen)
	}

	ebitenutil.DebugPrint(screen, fmt.Sprintf("FPS: %d", int(ebiten.ActualFPS())))
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("TPS: %d", int(ebiten.ActualTPS())), 0, 10)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideHeight
}

func (g *Game) newFood() {
	newFoodPos := vector.Vector{
		X: float64(rand.Intn(config.ScreenWidth/config.TileSize) * config.TileSize),
		Y: float64(rand.Intn(config.ScreenHeight/config.TileSize) * config.TileSize),
	}
	for _, wall := range g.walls {
		if wall.IsCollides(newFoodPos) {
			g.newFood()
			return
		}
	}
	s := g.snake
	for s != nil {
		if s.Pos == newFoodPos {
			fmt.Println("collision")
			newFoodPos = vector.Vector{
				X: float64(rand.Intn(config.ScreenWidth/config.TileSize) * config.TileSize),
				Y: float64(rand.Intn(config.ScreenHeight/config.TileSize) * config.TileSize),
			}
		}
		s = s.Next
	}
	g.foodPos = newFoodPos
}
