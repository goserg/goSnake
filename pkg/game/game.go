package game

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"goSnake/pkg/config"
	"goSnake/pkg/input"
	snakeField "goSnake/pkg/snake_field"
	"goSnake/pkg/text"
	"golang.org/x/image/colornames"
	"time"
)

type Game struct {
	snakeField *snakeField.SnakeField
}

func New() *Game {
	var g Game

	g.snakeField = snakeField.New()
	g.snakeField.DeathCallback = g.OnDeath

	return &g
}

func (g *Game) Update() error {
	input.Update()
	text.Update()

	if err := g.snakeField.Update(); err != nil {
		return err
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	var snakeFieldDrawingOptions ebiten.DrawImageOptions
	snakeFieldDrawingOptions.GeoM.Translate(config.FieldLeft, config.FieldTop)
	g.snakeField.Draw(screen, &snakeFieldDrawingOptions)

	ebitenutil.DebugPrint(screen, fmt.Sprintf("FPS: %d", int(ebiten.ActualFPS())))
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("TPS: %d", int(ebiten.ActualTPS())), 0, 10)

	text.Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideHeight
}

func (g *Game) OnDeath() {
	g.snakeField = snakeField.New()
	g.snakeField.DeathCallback = g.OnDeath

	text.New("YOU DIED", 200, 200,
		text.WithColor(colornames.Red),
		text.WithSize(100),
		text.WithFadeout(),
		text.WithMove(0, -0.5),
		text.WithLifespan(time.Second),
	)
}
