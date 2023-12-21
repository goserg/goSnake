package game

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"goSnake/pkg/image_manager"
	"goSnake/pkg/input"
	"goSnake/pkg/snake"
	"goSnake/pkg/utils/vector"
	"time"
)

const tileSize = 32

type Game struct {
	direction int
	snake     *snake.Snake

	foodPos vector.Vector

	mainTicker *time.Ticker
}

func New() *Game {
	var g Game
	g.snake = snake.New(vector.Vector{
		X: 32,
		Y: 32,
	})

	g.mainTicker = time.NewTicker(time.Second / 10)
	g.foodPos = vector.Vector{
		X: tileSize * 10,
		Y: tileSize * 10,
	}
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

		g.snake.Move(newPos)
		if g.snake.Pos == g.foodPos {
			g.snake.Grow()
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

	ebitenutil.DebugPrint(screen, fmt.Sprintf("FPS: %d", int(ebiten.ActualFPS())))
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("TPS: %d", int(ebiten.ActualTPS())), 0, 10)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideHeight
}
