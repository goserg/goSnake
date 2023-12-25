package resource

import (
	"github.com/hajimehoshi/ebiten/v2"
	"image/png"
)

type ImageTag int

const (
	ImageSkeleton ImageTag = iota + 1
	ImagePig
	ImageSword
)

var imageMap = map[ImageTag]string{
	ImageSkeleton: "74x17.png",
	ImagePig:      "94x16.png",
	ImageSword:    "28x36.png",
}

var imageCache map[ImageTag]*ebiten.Image

func init() {
	imageCache = make(map[ImageTag]*ebiten.Image)
}

func Image(tag ImageTag) *ebiten.Image {
	if img, ok := imageCache[tag]; ok {
		return img
	}
	fileName, ok := imageMap[tag]
	if !ok {
		panic("image tag unknown")
	}
	f, err := Images.Open("images/" + fileName)
	if err != nil {
		panic(err.Error())
	}
	img, err := png.Decode(f)
	if err != nil {
		panic(err.Error())
	}
	eimg := ebiten.NewImageFromImage(img)
	imageCache[tag] = eimg
	return eimg
}
