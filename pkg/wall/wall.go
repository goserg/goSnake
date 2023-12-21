package wall

import (
	"github.com/hajimehoshi/ebiten/v2"
	"goSnake/pkg/config"
	"goSnake/pkg/utils/vector"
)

type Wall struct {
	segments []*Segment
}

func New(pos vector.Vector) *Wall {
	var wall Wall

	wall.segments = []*Segment{
		NewSegment(vector.Vector{
			X: pos.X,
			Y: pos.Y,
		}),
		NewSegment(vector.Vector{
			X: pos.X + 1*config.TileSize,
			Y: pos.Y,
		}),
		NewSegment(vector.Vector{
			X: pos.X + 2*config.TileSize,
			Y: pos.Y,
		}),
	}
	return &wall
}

func (w *Wall) Draw(screen *ebiten.Image) {
	for _, segment := range w.segments {
		segment.Draw(screen)
	}
}

func (w *Wall) IsCollides(pos vector.Vector) bool {
	for _, segment := range w.segments {
		if segment.Pos == pos {
			return true
		}
	}
	return false
}
