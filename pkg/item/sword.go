package item

import (
	"goSnake/pkg/utils/vector"
	"goSnake/resource"
)

func NewSword(occupiedPositions map[vector.Vector]struct{}) *Item {
	var sword Item
	sword.findPosition(occupiedPositions)
	sword.Type = TypeSword
	sword.img = resource.Image(resource.ImageSword)
	return &sword
}
