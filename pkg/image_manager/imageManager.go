package image_manager

import (
	"github.com/hajimehoshi/ebiten/v2"
	"image/color"
)

var manager map[string]*ebiten.Image

func init() {
	manager = make(map[string]*ebiten.Image)
}

func Snake() *ebiten.Image {
	if img, ok := manager["snake"]; ok {
		return img
	}
	img := ebiten.NewImage(32, 32)
	img.Fill(color.RGBA{
		R: 200,
		G: 200,
		B: 0,
		A: 0,
	})
	manager["snake"] = img
	return img
}

func Food() *ebiten.Image {
	if img, ok := manager["food"]; ok {
		return img
	}
	img := ebiten.NewImage(32, 32)
	img.Fill(color.RGBA{
		R: 200,
		G: 0,
		B: 0,
		A: 0,
	})
	manager["food"] = img
	return img
}
