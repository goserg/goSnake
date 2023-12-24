package snake_field

import (
	"github.com/hajimehoshi/ebiten/v2"
	"goSnake/pkg/config"
	"goSnake/pkg/engine/signal"
	"goSnake/pkg/image_manager"
	"goSnake/pkg/input"
	"goSnake/pkg/item"
	"goSnake/pkg/snake"
	"goSnake/pkg/utils/vector"
	"image/color"
	"time"
)

const (
	No = iota
	Right
	Down
	Left
	Up
)

type SnakeField struct {
	isRunning bool

	speed int

	grid            *ebiten.Image
	queuedDirection int
	direction       int
	snake           *snake.Snake

	food *item.Item

	mainTicker *time.Ticker

	input *input.Handler

	EventEat   signal.Event[EventEatData]
	EventDeath signal.Event[EventSnakeDeathData]
}

type EventEatData struct {
	Name string
}

type EventSnakeDeathData struct {
}

func New(inputHandler *input.Handler) *SnakeField {
	var snakeField SnakeField

	snakeField.input = inputHandler

	snakeField.speed = 10
	snakeField.grid = image_manager.Grid()
	snakeField.snake = snake.New(vector.Vector{
		X: 32 * 5,
		Y: 32 * 5,
	})
	snakeField.mainTicker = time.NewTicker(calcTick(snakeField.speed))
	occupiedPositions := make(map[vector.Vector]struct{})
	for _, v := range snakeField.snake.Positions() {
		occupiedPositions[v] = struct{}{}
	}
	snakeField.food = item.NewFood(occupiedPositions)
	return &snakeField
}

func (sf *SnakeField) Update() error {
	if !sf.isRunning {
		return nil
	}
	switch sf.direction {
	case Left, Right:
		if sf.input.IsActionJustPressed(input.ActionUp) {
			sf.queuedDirection = Up
			break
		}
		if sf.input.IsActionJustPressed(input.ActionDown) {
			sf.queuedDirection = Down
			break
		}
	case Up, Down:
		if sf.input.IsActionJustPressed(input.ActionLeft) {
			sf.queuedDirection = Left
			break
		}
		if sf.input.IsActionJustPressed(input.ActionRight) {
			sf.queuedDirection = Right
			break
		}
	}

	select {
	case <-sf.mainTicker.C:
		if sf.queuedDirection != No {
			sf.direction = sf.queuedDirection
		}
		if sf.direction == No {
			sf.direction = Right
		}
		var newPos vector.Vector
		switch sf.direction {
		case Left:
			newPos = vector.Vector{
				X: sf.snake.Pos.X - 32,
				Y: sf.snake.Pos.Y,
			}
		case Right:
			newPos = vector.Vector{
				X: sf.snake.Pos.X + 32,
				Y: sf.snake.Pos.Y,
			}
		case Up:
			newPos = vector.Vector{
				X: sf.snake.Pos.X,
				Y: sf.snake.Pos.Y - 32,
			}
		case Down:
			newPos = vector.Vector{
				X: sf.snake.Pos.X,
				Y: sf.snake.Pos.Y + 32,
			}
		}
		switch {
		case newPos.X < 0:
			newPos.X = config.FieldWidth - config.TileSize
		case newPos.Y < 0:
			newPos.Y = config.FieldHeight - config.TileSize
		case newPos.X+config.TileSize > config.FieldWidth:
			newPos.X = 0
		case newPos.Y+config.TileSize > config.FieldHeight:
			newPos.Y = 0
		}

		s := sf.snake.Next
		for s != nil {
			if s.Pos == newPos {
				sf.EventDeath.Emit(EventSnakeDeathData{})
				return nil
			}
			s = s.Next
		}

		sf.snake.Move(newPos)

		occupiedPositions := make(map[vector.Vector]struct{})
		for _, v := range sf.snake.Positions() {
			occupiedPositions[v] = struct{}{}
		}

		if sf.snake.HeadPos() == sf.food.Pos() {
			sf.EventEat.Emit(EventEatData{Name: "lol kek"})
			sf.snake.Grow()
			sf.food = item.NewFood(occupiedPositions)
		}
	default:
	}
	return nil
}

func (sf *SnakeField) Draw(screen *ebiten.Image, opts *ebiten.DrawImageOptions) {
	snakeField := ebiten.NewImage(config.FieldWidth+1, config.FieldHeight+1)
	snakeField.Fill(color.RGBA{R: 10, G: 10, B: 10})

	snakeField.DrawImage(sf.grid, nil)
	if sf.snake != nil {
		sf.snake.Draw(snakeField, calcTick(sf.speed))
	}
	sf.food.Draw(snakeField)

	screen.DrawImage(snakeField, opts)
}

func (sf *SnakeField) SpeedUp() {
	sf.speed++
	sf.mainTicker = time.NewTicker(calcTick(sf.speed))
}

func (sf *SnakeField) GrowSnake() {
	sf.snake.Grow()
}

func (sf *SnakeField) Start() {
	sf.isRunning = true
}

func calcTick(speed int) time.Duration {
	return (time.Second) / time.Duration(speed)
}
