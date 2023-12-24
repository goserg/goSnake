package snake

import (
	"github.com/hajimehoshi/ebiten/v2"
	"goSnake/pkg/config"
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
	snake.Image = image_manager.Snake()
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
	nowTime := time.Now().Sub(s.MoveStarted)
	movedPart := float64(nowTime) / float64(tick)

	prevPos := s.PrevPos
	pos := s.Pos
	switch {
	case prevPos.X == 0 && s.Pos.X == config.FieldWidth-config.TileSize:
		pos.X = -config.TileSize
		var visualPos vector.Vector
		visualPos.X = prevPos.X + (pos.X-prevPos.X)*movedPart
		visualPos.Y = prevPos.Y + (pos.Y-prevPos.Y)*movedPart

		var snakeDrawOptions ebiten.DrawImageOptions
		snakeDrawOptions.GeoM.Translate(visualPos.X, visualPos.Y)
		screen.DrawImage(s.Image, &snakeDrawOptions)

		prevPos.X = config.FieldWidth
	case prevPos.X == config.FieldWidth-config.TileSize && s.Pos.X == 0:
		pos.X = config.FieldWidth
		var visualPos vector.Vector
		visualPos.X = prevPos.X + (pos.X-prevPos.X)*movedPart
		visualPos.Y = prevPos.Y + (pos.Y-prevPos.Y)*movedPart

		var snakeDrawOptions ebiten.DrawImageOptions
		snakeDrawOptions.GeoM.Translate(visualPos.X, visualPos.Y)
		screen.DrawImage(s.Image, &snakeDrawOptions)

		prevPos.X = -config.TileSize
	case prevPos.Y == 0 && s.Pos.Y == config.FieldHeight-config.TileSize:
		pos.Y = -config.TileSize
		var visualPos vector.Vector
		visualPos.X = prevPos.X + (pos.X-prevPos.X)*movedPart
		visualPos.Y = prevPos.Y + (pos.Y-prevPos.Y)*movedPart

		var snakeDrawOptions ebiten.DrawImageOptions
		snakeDrawOptions.GeoM.Translate(visualPos.X, visualPos.Y)
		screen.DrawImage(s.Image, &snakeDrawOptions)

		prevPos.Y = config.FieldHeight
	case prevPos.Y == config.FieldHeight-config.TileSize && s.Pos.Y == 0:
		pos.Y = config.FieldHeight
		var visualPos vector.Vector
		visualPos.X = prevPos.X + (pos.X-prevPos.X)*movedPart
		visualPos.Y = prevPos.Y + (pos.Y-prevPos.Y)*movedPart

		var snakeDrawOptions ebiten.DrawImageOptions
		snakeDrawOptions.GeoM.Translate(visualPos.X, visualPos.Y)
		screen.DrawImage(s.Image, &snakeDrawOptions)

		prevPos.Y = -config.TileSize
	}

	var visualPos vector.Vector
	visualPos.X = prevPos.X + (s.Pos.X-prevPos.X)*movedPart
	visualPos.Y = prevPos.Y + (s.Pos.Y-prevPos.Y)*movedPart

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

func (s *Snake) HeadPos() vector.Vector {
	return s.Pos
}

func (s *Snake) Positions() []vector.Vector {
	pos := []vector.Vector{s.Pos}
	segment := s.Next
	for segment != nil {
		pos = append(pos, segment.Pos)
		segment = segment.Next
	}
	return pos
}
