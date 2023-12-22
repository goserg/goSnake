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
	case prevPos.X == 0 && s.Pos.X == config.ScreenWidth-config.TileSize:
		pos.X = -config.TileSize
		var visualPos vector.Vector
		visualPos.X = prevPos.X + (pos.X-prevPos.X)*movedPart
		visualPos.Y = prevPos.Y + (pos.Y-prevPos.Y)*movedPart

		var snakeDrawOptions ebiten.DrawImageOptions
		snakeDrawOptions.GeoM.Translate(visualPos.X, visualPos.Y)
		screen.DrawImage(s.Image, &snakeDrawOptions)

		prevPos.X = config.ScreenWidth
	case prevPos.X == config.ScreenWidth-config.TileSize && s.Pos.X == 0:
		pos.X = config.ScreenWidth
		var visualPos vector.Vector
		visualPos.X = prevPos.X + (pos.X-prevPos.X)*movedPart
		visualPos.Y = prevPos.Y + (pos.Y-prevPos.Y)*movedPart

		var snakeDrawOptions ebiten.DrawImageOptions
		snakeDrawOptions.GeoM.Translate(visualPos.X, visualPos.Y)
		screen.DrawImage(s.Image, &snakeDrawOptions)

		prevPos.X = -config.TileSize
	case prevPos.Y == 0 && s.Pos.Y == config.ScreenHeight-config.TileSize:
		pos.Y = -config.TileSize
		var visualPos vector.Vector
		visualPos.X = prevPos.X + (pos.X-prevPos.X)*movedPart
		visualPos.Y = prevPos.Y + (pos.Y-prevPos.Y)*movedPart

		var snakeDrawOptions ebiten.DrawImageOptions
		snakeDrawOptions.GeoM.Translate(visualPos.X, visualPos.Y)
		screen.DrawImage(s.Image, &snakeDrawOptions)

		prevPos.Y = config.ScreenHeight
	case prevPos.Y == config.ScreenHeight-config.TileSize && s.Pos.Y == 0:
		pos.Y = config.ScreenHeight
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
