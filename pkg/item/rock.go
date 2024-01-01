package item

import (
	"github.com/goserg/zg/vector"
	"goSnake/resource"
)

func NewRock(occupiedPositions map[vector.Vector]struct{}) *Item {
	var sword Item
	sword.findPosition(occupiedPositions)
	sword.Type = TypeRock
	sword.img = resource.Image(resource.ImageRock)
	return &sword
}
