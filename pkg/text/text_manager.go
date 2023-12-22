package text

import (
	"github.com/google/uuid"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/examples/resources/fonts"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
	"log"
)

type tm struct {
	font  *opentype.Font
	fonts map[int]font.Face

	texts map[uuid.UUID]*textHolder
}

const dpi = 72

var textManager tm

func init() {
	tt, err := opentype.Parse(fonts.MPlus1pRegular_ttf)
	if err != nil {
		log.Fatal(err)
	}

	textManager = tm{
		font:  tt,
		fonts: make(map[int]font.Face),
		texts: make(map[uuid.UUID]*textHolder),
	}
}

func Draw(screen *ebiten.Image) {
	for _, holder := range textManager.texts {
		holder.Draw(screen)
	}
}

func Update() {
	var toDelete []uuid.UUID
	for id, holder := range textManager.texts {
		if holder.Update() {
			toDelete = append(toDelete, id)
		}
	}

	for _, id := range toDelete {
		delete(textManager.texts, id)
	}
}

func getFont(size int) font.Face {
	if f, ok := textManager.fonts[size]; ok {
		return f
	}
	mplusNormalFont, err := opentype.NewFace(textManager.font, &opentype.FaceOptions{
		Size:    float64(size),
		DPI:     dpi,
		Hinting: font.HintingVertical,
	})
	if err != nil {
		log.Fatal(err.Error())
	}
	textManager.fonts[size] = mplusNormalFont
	return mplusNormalFont
}
