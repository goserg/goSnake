package main

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"goSnake/pkg/config"
	"goSnake/pkg/game"
	"os"
)

func main() {
	ebiten.SetWindowSize(config.ScreenWidth, config.ScreenHeight)
	ebiten.SetWindowTitle("goSnake")
	if err := run(); err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}

func run() error {
	return ebiten.RunGame(game.New())
}
