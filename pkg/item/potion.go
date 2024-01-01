package item

import (
	"github.com/goserg/zg/vector"
	"goSnake/resource"

	"github.com/hajimehoshi/ebiten/v2"
)

func NewPotionSpeedUp(occupiedPositions map[vector.Vector]struct{}) *Item {
	var potion Item
	potion.findPosition(occupiedPositions)

	potionImg := ebiten.NewImageFromImage(resource.Image(resource.ImagePotionBlue))
	potionImg.DrawImage(resource.Image(resource.ImageSpeedUp), nil)
	potion.img = potionImg
	potion.Type = TypePotionSpeedUp

	return &potion
}
