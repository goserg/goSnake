package text

import (
	"github.com/google/uuid"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"goSnake/pkg/config"
	"image/color"
	"time"
)

type textHolder struct {
	id          uuid.UUID
	image       *ebiten.Image
	opts        *ebiten.DrawImageOptions
	destroyTime time.Time
	animations  []animation
}

type textOpts struct {
	size       int
	color      color.Color
	lifeSpan   time.Duration
	animations []animation
}

var zeroTime = time.Time{}

var defaultTextOpts = textOpts{
	size:       14,
	color:      color.White,
	lifeSpan:   0,
	animations: []animation{},
}

type textOption func(opts *textOpts)

func WithSize(size int) textOption {
	return func(opts *textOpts) {
		opts.size = size
	}
}

func WithColor(color color.Color) textOption {
	return func(opts *textOpts) {
		opts.color = color
	}
}

func WithLifespan(lifespan time.Duration) textOption {
	return func(opts *textOpts) {
		opts.lifeSpan = lifespan
	}
}

func WithFadeout() textOption {
	return func(opts *textOpts) {
		opts.animations = append(opts.animations, fadeout())
	}
}

func WithMove(dx float64, dy float64) textOption {
	return func(opts *textOpts) {
		opts.animations = append(opts.animations, move(dx, dy))
	}
}

func New(str string, posX float64, posY float64, opts ...textOption) {
	o := defaultTextOpts
	for _, opt := range opts {
		opt(&o)
	}
	var holder textHolder
	img := ebiten.NewImage(config.ScreenWidth, config.ScreenHeight)
	text.Draw(img, str, getFont(o.size), 0, o.size, o.color)

	holder.image = img
	holder.opts = &ebiten.DrawImageOptions{}
	holder.id = uuid.New()
	holder.animations = o.animations
	if o.lifeSpan > 0 {
		holder.destroyTime = time.Now().Add(o.lifeSpan)
	}
	holder.opts.GeoM.Translate(posX, posY)

	textManager.texts[holder.id] = &holder
}

func (t *textHolder) Draw(screen *ebiten.Image) {
	screen.DrawImage(t.image, t.opts)
}

func (t *textHolder) Update() (deleteMe bool) {
	if t.destroyTime != zeroTime && t.destroyTime.Before(time.Now()) {
		return true
	}
	for _, animation := range t.animations {
		animation(t.opts)
	}
	return false
}
