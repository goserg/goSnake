package config

import "time"

const (
	TileSize = 32

	ScreenWidth  = TileSize * 16 * 2
	ScreenHeight = TileSize * 9 * 2

	Tick = time.Second / 1
)
