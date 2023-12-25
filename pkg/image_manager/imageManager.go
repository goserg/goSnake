package image_manager

import (
	"github.com/hajimehoshi/ebiten/v2"
	"goSnake/pkg/config"
	"goSnake/resource"
	"image/color"
	"image/png"
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
		R: 210,
		G: 210,
		B: 4,
		A: 255,
	})
	for x := 0; x < img.Bounds().Dx(); x++ {
		img.Set(x, 0, color.RGBA{
			R: 255,
			G: 255,
			B: 0,
			A: 255,
		})
		img.Set(x, img.Bounds().Dy()-1, color.RGBA{
			R: 255,
			G: 255,
			B: 0,
			A: 255,
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
		R: 204,
		G: 4,
		B: 4,
		A: 255,
	})
	manager["wall"] = img
	return img
}

func Grid() *ebiten.Image {
	if img, ok := manager["grid"]; ok {
		return img
	}
	img := ebiten.NewImage(config.ScreenWidth, config.ScreenHeight)

	for x := 0; x < config.FieldWidth+config.TileSize; x += config.TileSize {
		for y := 0; y < config.FieldHeight; y++ {
			img.Set(x, y, color.RGBA{
				R: 20,
				G: 20,
				B: 20,
				A: 0,
			})
		}
	}
	for y := 0; y < config.FieldHeight+config.TileSize; y += config.TileSize {
		for x := 0; x < config.FieldWidth; x++ {
			img.Set(x, y, color.RGBA{
				R: 20,
				G: 20,
				B: 20,
				A: 0,
			})
		}
	}
	manager["grid"] = img
	return img
}

func Sword() *ebiten.Image {
	if img, ok := manager["sword"]; ok {
		return img
	}
	f, err := resource.Images.Open("images/28x36.png")
	if err != nil {
		return nil
	}
	defer f.Close()

	img, err := png.Decode(f)
	if err != nil {
		panic(err.Error())
	}
	eimg := ebiten.NewImageFromImage(img)
	manager["sword"] = eimg
	return eimg
}

func Pig() *ebiten.Image {
	if img, ok := manager["pig"]; ok {
		return img
	}
	f, err := resource.Images.Open("images/94x16.png")
	if err != nil {
		return nil
	}
	defer f.Close()

	img, err := png.Decode(f)
	if err != nil {
		panic(err.Error())
	}
	eimg := ebiten.NewImageFromImage(img)
	manager["pig"] = eimg
	return eimg
}
