package snake

import (
	"github.com/hajimehoshi/ebiten/v2"
	"goSnake/pkg/image_manager"
	"goSnake/pkg/utils/vector"
)

type Snake struct {
	Pos   vector.Vector
	Image *ebiten.Image
	Next  *Snake
}

func New() *Snake {
	var snake Snake

	snake.Image = image_manager.Snake()
	return &snake
}

func (s *Snake) Draw(screen *ebiten.Image) {
	var snakeDrawOptions ebiten.DrawImageOptions
	snakeDrawOptions.GeoM.Translate(s.Pos.X, s.Pos.Y)
	screen.DrawImage(s.Image, &snakeDrawOptions)

	if s.Next != nil {
		s.Next.Draw(screen)
	}
}
