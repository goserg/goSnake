package input

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"time"
)

type Handler struct {
	id     int
	keymap keymap
}

func NewHandler() *Handler {
	var handler Handler
	handler.keymap = keyboardKeymap
	for action, keys := range gamepadKeymap {
		handler.keymap[action] = append(handler.keymap[action], keys...)
	}
	return &handler
}

func (h *Handler) Update(delta time.Duration) {
	if h.id == 0 {
		for _, g := range ebiten.AppendGamepadIDs([]ebiten.GamepadID{}) {
			h.id = int(g)
		}
	}
	//pressed := inpututil.AppendJustPressedGamepadButtons(ebiten.GamepadID(h.id), []ebiten.GamepadButton{})
	//if len(pressed) != 0 {
	//	for _, keys := range h.keymap {
	//		for _, key := range keys {
	//			for _, button := range pressed {
	//				if int(button) == key.code {
	//					fmt.Println(key.name)
	//				}
	//			}
	//		}
	//	}
	//}
}

func (h *Handler) IsActionJustPressed(action Action) bool {
	keys, ok := h.keymap[action]
	if !ok {
		// действие не замаплено
		return false
	}
	for _, key := range keys {
		switch key.kind {
		case keyKeyboard:
			if inpututil.IsKeyJustPressed(ebiten.Key(key.code)) {
				return true
			}
		case keyGamepad:
			if inpututil.IsGamepadButtonJustPressed(ebiten.GamepadID(h.id), ebiten.GamepadButton(key.code)) {
				return true
			}
		}
	}
	return false
}
