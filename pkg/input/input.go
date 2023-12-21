package input

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

const (
	No = iota
	Right
	Down
	Left
	Up
)

type manager struct {
	keysPressed []ebiten.Key
}

var m manager

func Update() {
	m.keysPressed = inpututil.AppendJustPressedKeys([]ebiten.Key{})
}

func IsLeftPressed() bool {
	for _, key := range m.keysPressed {
		if ebiten.KeyLeft == key {
			return true
		}
	}
	return false
}
func IsRightPressed() bool {
	for _, key := range m.keysPressed {
		if ebiten.KeyRight == key {
			return true
		}
	}
	return false
}
func IsUpPressed() bool {
	for _, key := range m.keysPressed {
		if ebiten.KeyUp == key {
			return true
		}
	}
	return false
}
func IsDownPressed() bool {
	for _, key := range m.keysPressed {
		if ebiten.KeyDown == key {
			return true
		}
	}
	return false
}

func DirectionH() int {
	var dir int
	if IsLeftPressed() {
		dir -= 1
	}
	if IsRightPressed() {
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
	if IsUpPressed() {
		dir -= 1
	}
	if IsDownPressed() {
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
