package wall

import (
	"github.com/hajimehoshi/ebiten/v2"
	"goSnake/pkg/image_manager"
	"goSnake/pkg/utils/vector"
)

type Segment struct {
	Image *ebiten.Image
	Pos   vector.Vector
}

func NewSegment(pos vector.Vector) *Segment {
	var segment Segment

	segment.Pos = pos
	segment.Image = image_manager.Wall()

	return &segment
}

func (s *Segment) Draw(screen *ebiten.Image) {
	var wallDrawingOptions ebiten.DrawImageOptions
	wallDrawingOptions.GeoM.Translate(s.Pos.X, s.Pos.Y)

	screen.DrawImage(s.Image, &wallDrawingOptions)
}
