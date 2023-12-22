package input

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"log"
)

const (
	No = iota
	Right
	Down
	Left
	Up
)

type manager struct {
	keysPressed    []ebiten.Key
	buttonsPressed []ebiten.GamepadButton
	gamepadIDsBuf  []ebiten.GamepadID
	gamepadIDs     map[ebiten.GamepadID]struct{}
}

func init() {
	m = manager{gamepadIDs: map[ebiten.GamepadID]struct{}{}}
}

var m manager

func Update() {
	m.keysPressed = inpututil.AppendJustPressedKeys([]ebiten.Key{})

	// Log the gamepad connection events.
	m.gamepadIDsBuf = inpututil.AppendJustConnectedGamepadIDs(m.gamepadIDsBuf[:0])
	for _, id := range m.gamepadIDsBuf {
		log.Printf("gamepad connected: id: %d, SDL ID: %s", id, ebiten.GamepadSDLID(id))
		m.gamepadIDs[id] = struct{}{}
	}
	for id := range m.gamepadIDs {
		if inpututil.IsGamepadJustDisconnected(id) {
			log.Printf("gamepad disconnected: id: %d", id)
			delete(m.gamepadIDs, id)
		}
	}
	m.buttonsPressed = []ebiten.GamepadButton{}
	for id := range m.gamepadIDs {
		pressed := inpututil.AppendJustPressedGamepadButtons(id, []ebiten.GamepadButton{})
		for _, button := range pressed {
			m.buttonsPressed = append(m.buttonsPressed, button)
		}
	}
}

func IsLeftJustPressed() bool {
	for _, key := range m.keysPressed {
		if ebiten.KeyLeft == key {
			return true
		}
	}
	for _, button := range m.buttonsPressed {
		if ebiten.GamepadButton13 == button {
			return true
		}
	}

	return false
}
func IsRightJustPressed() bool {
	for _, key := range m.keysPressed {
		if ebiten.KeyRight == key {
			return true
		}
	}
	for _, button := range m.buttonsPressed {
		if ebiten.GamepadButton11 == button {
			return true
		}
	}
	return false
}
func IsUpJustPressed() bool {
	for _, key := range m.keysPressed {
		if ebiten.KeyUp == key {
			return true
		}
	}
	for _, button := range m.buttonsPressed {
		if ebiten.GamepadButton10 == button {
			return true
		}
	}
	return false
}
func IsDownJustPressed() bool {
	for _, key := range m.keysPressed {
		if ebiten.KeyDown == key {
			return true
		}
	}
	for _, button := range m.buttonsPressed {
		if ebiten.GamepadButton12 == button {
			return true
		}
	}
	return false
}

func DirectionH() int {
	var dir int
	if IsLeftJustPressed() {
		dir -= 1
	}
	if IsRightJustPressed() {
		dir += 1
	}
	if dir == 1 {
		return Right
	}
	if dir == -1 {
		return Left
	}
	return No
}

func DirectionV() int {
	var dir int
	if IsUpJustPressed() {
		dir -= 1
	}
	if IsDownJustPressed() {
		dir += 1
	}
	if dir == 1 {
		return Down
	}
	if dir == -1 {
		return Up
	}
	return No
}

func IsF1Pressed() bool {
	return inpututil.IsKeyJustPressed(ebiten.KeyF1)
}
