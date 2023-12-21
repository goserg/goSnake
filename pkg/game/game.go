package game

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"goSnake/pkg/image_manager"
	"goSnake/pkg/input"
	"time"
)

const tileSize = 32

type Game struct {
	snakePosX int
	snakePosY int
	direction int

	foodPosX int
	foodPosY int

	mainTicker *time.Ticker
}

func New() *Game {
	var g Game

	g.mainTicker = time.NewTicker(time.Second / 2)
	g.foodPosX = tileSize * 10
	g.foodPosY = tileSize * 10
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
			g.snakePosX -= tileSize
		case input.Right:
			g.snakePosX += tileSize
		case input.Up:
			g.snakePosY -= tileSize
		case input.Down:
			g.snakePosY += tileSize
		}
	default:
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	var snakeDrawOptions ebiten.DrawImageOptions
	snakeDrawOptions.GeoM.Translate(float64(g.snakePosX), float64(g.snakePosY))
	screen.DrawImage(image_manager.Snake(), &snakeDrawOptions)

	var foodDrawOptions ebiten.DrawImageOptions
	foodDrawOptions.GeoM.Translate(float64(g.foodPosX), float64(g.foodPosY))
	screen.DrawImage(image_manager.Food(), &foodDrawOptions)

	ebitenutil.DebugPrint(screen, fmt.Sprintf("FPS: %d", int(ebiten.ActualFPS())))
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("TPS: %d", int(ebiten.ActualTPS())), 0, 10)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideHeight
}
