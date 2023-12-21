package snake

import (
	"github.com/hajimehoshi/ebiten/v2"
	"goSnake/pkg/image_manager"
	"goSnake/pkg/utils/vector"
)

type Snake struct {
	direction int
	Pos       vector.Vector
	Image     *ebiten.Image
	Next      *Snake
}

func New(pos vector.Vector) *Snake {
	var snake Snake
	snake.Pos = pos
	snake.Image = image_manager.SnakeSingle()
	return &snake
}

func (s *Snake) Grow() {
	if s.Next != nil {
		s.Next.Grow()
		return
	}
	s.Next = New(s.Pos)
}

func (s *Snake) Draw(screen *ebiten.Image) {
	var snakeDrawOptions ebiten.DrawImageOptions
	snakeDrawOptions.GeoM.Translate(s.Pos.X, s.Pos.Y)
	screen.DrawImage(s.Image, &snakeDrawOptions)

	if s.Next != nil {
		s.Next.Draw(screen)
	}
}

func (s *Snake) Move(pos vector.Vector) {
	oldpos := s.Pos
	s.Pos = pos
	if s.Next != nil {
		s.Next.Move(oldpos)
	}
}
