package main

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"goSnake/pkg/game"
	"os"
)

func main() {
	ebiten.SetWindowSize(32*16*2, 32*9*2)
	ebiten.SetWindowTitle("goSnake")
	if err := run(); err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}

func run() error {
	return ebiten.RunGame(game.New())
}
