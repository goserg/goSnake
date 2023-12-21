package game

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"goSnake/pkg/image_manager"
	"goSnake/pkg/input"
	"goSnake/pkg/utils/vector"
	"time"
)

const tileSize = 32

type Game struct {
	snakePos  vector.Vector
	direction int

	foodPos vector.Vector

	mainTicker *time.Ticker
}

func New() *Game {
	var g Game

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
		switch g.direction {
		case input.Left:
			g.snakePos.X -= tileSize
		case input.Right:
			g.snakePos.X += tileSize
		case input.Up:
			g.snakePos.Y -= tileSize
		case input.Down:
			g.snakePos.Y += tileSize
		}
	default:
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	var snakeDrawOptions ebiten.DrawImageOptions
	snakeDrawOptions.GeoM.Translate(g.snakePos.X, g.snakePos.Y)
	screen.DrawImage(image_manager.Snake(), &snakeDrawOptions)

	var foodDrawOptions ebiten.DrawImageOptions
	foodDrawOptions.GeoM.Translate(g.foodPos.X, g.foodPos.Y)
	screen.DrawImage(image_manager.Food(), &foodDrawOptions)

	ebitenutil.DebugPrint(screen, fmt.Sprintf("FPS: %d", int(ebiten.ActualFPS())))
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("TPS: %d", int(ebiten.ActualTPS())), 0, 10)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideHeight
}
