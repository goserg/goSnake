package item

import (
	"github.com/goserg/zg/vector"
	"goSnake/resource"
)

func NewBlood(pos vector.Vector[float64]) *Item {
	var blood Item
	blood.position = pos
	blood.img = resource.Image(resource.ImageBlood)
	return &blood
}
