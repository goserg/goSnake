package text

import "github.com/hajimehoshi/ebiten/v2"

type animation func(options *ebiten.DrawImageOptions)

func fadeout() animation {
	return func(options *ebiten.DrawImageOptions) {
		options.ColorScale.Scale(0.95, 0.95, 0.95, 1)
	}
}

func move(dx float64, dy float64) animation {
	return func(options *ebiten.DrawImageOptions) {
		options.GeoM.Translate(dx, dy)
	}
}
