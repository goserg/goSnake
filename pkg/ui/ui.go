package ui

import (
	"image/color"
	"log"
	"os"
	"time"

	"github.com/goserg/zg/signal"

	"github.com/ebitenui/ebitenui"
	"github.com/ebitenui/ebitenui/image"
	"github.com/ebitenui/ebitenui/widget"
	"github.com/golang/freetype/truetype"
	"github.com/hajimehoshi/ebiten/v2"
	"golang.org/x/image/font"
	"golang.org/x/image/font/gofont/goregular"
)

type UI struct {
	ui *ebitenui.UI

	EventStartPressed signal.Event[struct{}]
}

func New() *UI {
	var ui UI
	rootContainer := widget.NewContainer(
		// the container will use an anchor layout to layout its single child widget
		widget.ContainerOpts.Layout(
			widget.NewAnchorLayout(),
		),
	)
	eui := &ebitenui.UI{
		Container: rootContainer,
	}
	// This loads a font and creates a font face.
	ttfFont, err := truetype.Parse(goregular.TTF)
	if err != nil {
		log.Fatal("Error Parsing Font", err)
	}
	fontFace := truetype.NewFace(ttfFont, &truetype.Options{
		Size: 32,
	})

	menu := widget.NewContainer(
		// the container will use an anchor layout to layout its single child widget
		widget.ContainerOpts.Layout(
			widget.NewRowLayout(
				widget.RowLayoutOpts.Direction(widget.DirectionVertical),
				widget.RowLayoutOpts.Spacing(10),
			),
		),
		widget.ContainerOpts.WidgetOpts(
			widget.WidgetOpts.LayoutData(widget.AnchorLayoutData{
				HorizontalPosition: widget.AnchorLayoutPositionCenter,
				VerticalPosition:   widget.AnchorLayoutPositionCenter,
				StretchHorizontal:  false,
				StretchVertical:    false,
			})),
	)

	startButton := Button("press start", fontFace, func(args *widget.ButtonClickedEventArgs) {
		ui.EventStartPressed.Emit(struct{}{})
	})
	exitButton := Button("exit", fontFace, func(args *widget.ButtonClickedEventArgs) {
		os.Exit(0)
	})

	menu.AddChild(startButton)
	menu.AddChild(exitButton)

	rootContainer.AddChild(menu)

	ui.ui = eui
	return &ui
}

func Button(text string, fontFace font.Face, handler func(args *widget.ButtonClickedEventArgs)) *widget.Button {
	bimage, _ := loadButtonImage()
	return widget.NewButton(
		widget.ButtonOpts.WidgetOpts(
			// instruct the container's anchor layout to center the button both horizontally and vertically
			widget.WidgetOpts.LayoutData(widget.AnchorLayoutData{
				HorizontalPosition: widget.AnchorLayoutPositionCenter,
				VerticalPosition:   widget.AnchorLayoutPositionCenter,
			}),
		),
		widget.ButtonOpts.Text(text, fontFace, &widget.ButtonTextColor{
			Idle:     color.White,
			Disabled: color.Gray{Y: 50},
		}),
		widget.ButtonOpts.TextPadding(widget.Insets{10, 10, 10, 10}),
		widget.ButtonOpts.Image(bimage),
		widget.ButtonOpts.ClickedHandler(handler),
	)
}

func (ui *UI) Update(delta time.Duration) {
	ui.ui.Update()
}

func (ui *UI) Draw(screen *ebiten.Image) {
	ui.ui.Draw(screen)
}

func (ui *UI) HideMenu() {
	ui.ui.Container.GetWidget().Visibility = widget.Visibility_Hide
}

func (ui *UI) ToggleMenu() {
	switch ui.ui.Container.GetWidget().Visibility {
	case widget.Visibility_Hide:
		ui.ui.Container.GetWidget().Visibility = widget.Visibility_Show
	case widget.Visibility_Show:
		ui.ui.Container.GetWidget().Visibility = widget.Visibility_Hide

	}
}

func (ui *UI) ShowMenu() {
	ui.ui.Container.GetWidget().Visibility = widget.Visibility_Show
}

func loadButtonImage() (*widget.ButtonImage, error) {
	idle := image.NewNineSliceColor(color.Transparent)

	hover := image.NewNineSliceColor(color.RGBA{50, 50, 50, 10})

	pressed := image.NewNineSliceColor(color.Transparent)

	return &widget.ButtonImage{
		Idle:    idle,
		Hover:   hover,
		Pressed: pressed,
	}, nil
}
