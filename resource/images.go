package resource

import (
	"github.com/hajimehoshi/ebiten/v2"
	"image/png"
)

type ImageTag string

const (
	ImageSkeleton ImageTag = "skeleton"
	ImagePig               = "pig"
	ImageSword             = "sword"
	ImageRock              = "rock"
	ImageBlood             = "blood"
	ImagePlusOne           = "plus_one"
)

var imageMap = map[ImageTag]string{
	ImageSkeleton: "74x17.png",
	ImagePig:      "94x16.png",
	ImageSword:    "28x36.png",
	ImageRock:     "28x18.png",
	ImageBlood:    "51x23.png",
	ImagePlusOne:  "58x55.png",
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
