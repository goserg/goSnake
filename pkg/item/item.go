package item

import (
	"github.com/hajimehoshi/ebiten/v2"
	"goSnake/pkg/config"
	"goSnake/pkg/image_manager"
	"goSnake/pkg/utils/vector"
	"math/rand"
)

type Item struct {
	position vector.Vector
	img      *ebiten.Image
}

func NewFood(occupiedPositions map[vector.Vector]struct{}) *Item {
	var food Item
	for {
		food.position = vector.Vector{
			X: float64(rand.Intn(config.FieldWidth/config.TileSize) * config.TileSize),
			Y: float64(rand.Intn(config.FieldHeight/config.TileSize) * config.TileSize),
		}
		if _, ok := occupiedPositions[food.position]; !ok {
			break
		}
	}

	food.img = image_manager.Food()

	return &food
}

func (i *Item) Draw(screen *ebiten.Image) {
	var foodDrawOptions ebiten.DrawImageOptions
	foodDrawOptions.GeoM.Translate(i.position.X, i.position.Y)
	screen.DrawImage(image_manager.Food(), &foodDrawOptions)
}

func (i *Item) Pos() vector.Vector {
	return i.position
}
