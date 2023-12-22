package snake

import (
	"github.com/hajimehoshi/ebiten/v2"
	"goSnake/pkg/image_manager"
	"goSnake/pkg/utils/vector"
	"time"
)

type Snake struct {
	direction   int
	PrevPos     vector.Vector
	Pos         vector.Vector
	MoveStarted time.Time
	Image       *ebiten.Image
	Next        *Snake
}

func New(pos vector.Vector) *Snake {
	var snake Snake
	snake.Pos = pos
	snake.PrevPos = pos
	snake.Image = image_manager.SnakeSingle()
	return &snake
}

func (s *Snake) Grow() {
	if s.Next != nil {
		s.Next.Grow()
		return
	}
	s.Next = New(s.PrevPos)
}

func (s *Snake) Draw(screen *ebiten.Image, tick time.Duration) {
	var visualPos vector.Vector
	nowTime := time.Now().Sub(s.MoveStarted)
	movedPart := float64(nowTime) / float64(tick)
	visualPos.X = s.PrevPos.X + (s.Pos.X-s.PrevPos.X)*movedPart
	visualPos.Y = s.PrevPos.Y + (s.Pos.Y-s.PrevPos.Y)*movedPart

	var snakeDrawOptions ebiten.DrawImageOptions
	snakeDrawOptions.GeoM.Translate(visualPos.X, visualPos.Y)
	screen.DrawImage(s.Image, &snakeDrawOptions)

	if s.Next != nil {
		s.Next.Draw(screen, tick)
	}
}

func (s *Snake) Move(pos vector.Vector) {
	s.PrevPos = s.Pos
	s.Pos = pos
	s.MoveStarted = time.Now()
	if s.Next != nil {
		s.Next.Move(s.PrevPos)
	}
}
