package item

import (
	"github.com/goserg/zg/vector"
	"goSnake/resource"
)

func NewSword(occupiedPositions map[vector.Vector[float64]]struct{}) *Item {
	var sword Item
	sword.findPosition(occupiedPositions)
	sword.Type = TypeSword
	sword.img = resource.Image(resource.ImageSword)
	return &sword
}
