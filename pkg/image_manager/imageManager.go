package image_manager

import (
	"github.com/hajimehoshi/ebiten/v2"
	"goSnake/pkg/config"
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
	img := ebiten.NewImage(config.TileSize, config.TileSize)
	img.Fill(color.RGBA{
		R: 200,
		G: 200,
		B: 0,
		A: 0,
	})
	manager["snake"] = img
	return img
}

func SnakeSingle() *ebiten.Image {
	if img, ok := manager["snake"]; ok {
		return img
	}
	img := ebiten.NewImage(config.TileSize, config.TileSize)
	img.Fill(color.RGBA{
		R: 200,
		G: 200,
		B: 0,
		A: 0,
	})
	for x := 0; x < img.Bounds().Dx(); x++ {
		img.Set(x, 0, color.RGBA{
			R: 255,
			G: 255,
			B: 0,
			A: 0,
		})
		img.Set(x, img.Bounds().Dy()-1, color.RGBA{
			R: 255,
			G: 255,
			B: 0,
			A: 0,
		})
	}
	for y := 0; y < img.Bounds().Dy(); y++ {
		img.Set(0, y, color.RGBA{
			R: 255,
			G: 255,
			B: 0,
			A: 0,
		})
		img.Set(img.Bounds().Dx()-1, y, color.RGBA{
			R: 255,
			G: 255,
			B: 0,
			A: 0,
		})
	}
	manager["snake"] = img
	return img
}

func Food() *ebiten.Image {
	if img, ok := manager["food"]; ok {
		return img
	}
	img := ebiten.NewImage(config.TileSize, config.TileSize)
	img.Fill(color.RGBA{
		R: 0,
		G: 200,
		B: 0,
		A: 0,
	})
	manager["food"] = img
	return img
}

func Wall() *ebiten.Image {
	if img, ok := manager["wall"]; ok {
		return img
	}
	img := ebiten.NewImage(config.TileSize, config.TileSize)
	img.Fill(color.RGBA{
		R: 200,
		G: 0,
		B: 0,
		A: 0,
	})
	manager["wall"] = img
	return img
}
