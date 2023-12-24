package input

import "github.com/hajimehoshi/ebiten/v2"

type Action int

type Key struct {
	code int
	kind keyKind
	name string
}

type keymap map[Action][]Key

type keyKind int

const (
	keyKeyboard keyKind = iota + 1
	keyGamepad  keyKind = iota + 1
)

var gamepadKeymap = keymap{
	ActionLeft:  {Key{code: 13, kind: keyGamepad, name: "gamepad left"}},
	ActionRight: {Key{code: 11, kind: keyGamepad, name: "gamepad right"}},
	ActionUp:    {Key{code: 10, kind: keyGamepad, name: "gamepad up"}},
	ActionDown:  {Key{code: 12, kind: keyGamepad, name: "gamepad down"}},
	Debug:       {Key{code: 6, kind: keyGamepad, name: "gamepad debug"}},
	Pause:       {Key{code: 7, kind: keyGamepad, name: "gamepad pause"}},
}

var keyboardKeymap = keymap{
	ActionLeft: {
		Key{code: int(ebiten.KeyA), kind: keyKeyboard},
		Key{code: int(ebiten.KeyLeft), kind: keyKeyboard},
	},
	ActionRight: {
		Key{code: int(ebiten.KeyD), kind: keyKeyboard},
		Key{code: int(ebiten.KeyRight), kind: keyKeyboard},
	},
	ActionUp: {
		Key{code: int(ebiten.KeyW), kind: keyKeyboard},
		Key{code: int(ebiten.KeyUp), kind: keyKeyboard},
	},
	ActionDown: {
		Key{code: int(ebiten.KeyS), kind: keyKeyboard},
		Key{code: int(ebiten.KeyDown), kind: keyKeyboard},
	},
	Debug: {
		Key{
			code: int(ebiten.KeyF1),
			kind: keyKeyboard,
		},
	},
	Pause: {
		Key{
			code: int(ebiten.KeyEscape),
			kind: keyKeyboard,
		},
	},
}
