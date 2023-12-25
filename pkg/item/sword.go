package item

import (
	"goSnake/pkg/image_manager"
	"goSnake/pkg/utils/vector"
)

func NewSword(occupiedPositions map[vector.Vector]struct{}) *Item {
	var sword Item
	sword.findPosition(occupiedPositions)
	sword.Type = TypeSword
	sword.img = image_manager.Sword()
	return &sword
}
