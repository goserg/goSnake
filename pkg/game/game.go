package game

import "github.com/hajimehoshi/ebiten/v2"

type Game struct {
}

func New() *Game {
	var g Game

	return &g
}

func (g Game) Update() error {
	return nil
}

func (g Game) Draw(screen *ebiten.Image) {
}

func (g Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideHeight
}
