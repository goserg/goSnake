package item

import (
	"goSnake/pkg/utils/vector"
	"goSnake/resource"
)

func NewBlood(pos vector.Vector) *Item {
	var blood Item
	blood.position = pos
	blood.img = resource.Image(resource.ImageBlood)
	return &blood
}
