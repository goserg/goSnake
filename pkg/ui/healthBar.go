package ui

import (
	"github.com/ebitenui/ebitenui"
	"github.com/ebitenui/ebitenui/image"
	"github.com/ebitenui/ebitenui/widget"
	"github.com/hajimehoshi/ebiten/v2"
	"image/color"
)

type HealthBar struct {
	ui       *ebitenui.UI
	setValue func(int) bool
}

func NewHealthBar() *HealthBar {
	var hb HealthBar

	rootContainer := widget.NewContainer(
		widget.ContainerOpts.Layout(
			widget.NewAnchorLayout(),
		),
	)
	ui := ebitenui.UI{
		Container:           rootContainer,
		DisableDefaultFocus: false,
	}
	bar := widget.NewProgressBar(
		widget.ProgressBarOpts.WidgetOpts(
			widget.WidgetOpts.MinSize(64, 3)),
		widget.ProgressBarOpts.Images(
			&widget.ProgressBarImage{
				Idle:     image.NewNineSliceColor(color.RGBA{R: 50, G: 50, B: 50, A: 10}),
				Hover:    image.NewNineSliceColor(color.RGBA{R: 50, G: 50, B: 50, A: 10}),
				Disabled: image.NewNineSliceColor(color.RGBA{R: 50, G: 50, B: 50, A: 10}),
			},
			&widget.ProgressBarImage{
				Idle:     image.NewNineSliceColor(color.RGBA{R: 150, G: 150, B: 150, A: 150}),
				Hover:    image.NewNineSliceColor(color.RGBA{R: 150, G: 150, B: 150, A: 150}),
				Disabled: image.NewNineSliceColor(color.RGBA{R: 150, G: 150, B: 150, A: 150}),
			},
		),
	)
	bar.SetCurrent(100)
	rootContainer.AddChild(bar)

	hb.ui = &ui
	hb.setValue = bar.SetCurrent

	return &hb
}

func (hb *HealthBar) Draw(screen *ebiten.Image) {
	hb.ui.Draw(screen)
}

func (hb *HealthBar) Set(value int) {
	hb.setValue(value)
}
