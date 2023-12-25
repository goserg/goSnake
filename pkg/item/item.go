package item

import (
	"github.com/hajimehoshi/ebiten/v2"
	"goSnake/pkg/config"
	"goSnake/pkg/image_manager"
	"goSnake/pkg/utils/vector"
	"math/rand"
)

type Item struct {
	Type     Type
	position vector.Vector
	img      *ebiten.Image
}

func NewFood(occupiedPositions map[vector.Vector]struct{}) *Item {
	var food Item
	food.findPosition(occupiedPositions)

	food.img = image_manager.Food()
	food.Type = TypeFood

	return &food
}

func (i *Item) findPosition(occupiedPositions map[vector.Vector]struct{}) {
	var redoThisMethodCounter int
	for redoThisMethodCounter < 1000 {
		redoThisMethodCounter++
		i.position = vector.Vector{
			X: float64(rand.Intn(config.FieldWidth/config.TileSize) * config.TileSize),
			Y: float64(rand.Intn(config.FieldHeight/config.TileSize) * config.TileSize),
		}
		if _, ok := occupiedPositions[i.position]; !ok {
			return
		}
	}
	panic("redoThisMethodCounter reached maximum")
}

func (i *Item) Draw(screen *ebiten.Image) {
	var foodDrawOptions ebiten.DrawImageOptions
	foodDrawOptions.GeoM.Translate(i.position.X, i.position.Y)
	screen.DrawImage(i.img, &foodDrawOptions)
}

func (i *Item) Pos() vector.Vector {
	return i.position
}
