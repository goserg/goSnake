package snake_field

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"goSnake/pkg/config"
	"goSnake/pkg/image_manager"
	"goSnake/pkg/input"
	"goSnake/pkg/snake"
	"goSnake/pkg/utils/vector"
	"image/color"
	"math/rand"
	"time"
)

type SnakeField struct {
	speed int

	grid            *ebiten.Image
	queuedDirection int
	direction       int
	snake           *snake.Snake

	foodPos vector.Vector

	mainTicker *time.Ticker

	DeathCallback func()
	EatCallback   func()
}

func New() *SnakeField {
	var snakeField SnakeField

	snakeField.speed = 10
	snakeField.grid = image_manager.Grid()
	snakeField.snake = snake.New(vector.Vector{
		X: 32 * 5,
		Y: 32 * 5,
	})
	snakeField.mainTicker = time.NewTicker(calcTick(snakeField.speed))
	snakeField.newFood()
	return &snakeField
}

func (sf *SnakeField) Update() error {
	input.Update()

	switch sf.direction {
	case input.Left, input.Right:
		if input.DirectionV() != input.No {
			sf.queuedDirection = input.DirectionV()
		}
	case input.Up, input.Down:
		if input.DirectionH() != input.No {
			sf.queuedDirection = input.DirectionH()
		}
	}

	select {
	case <-sf.mainTicker.C:
		if sf.queuedDirection != input.No {
			sf.direction = sf.queuedDirection
		}
		if sf.direction == input.No {
			sf.direction = input.Right
		}
		var newPos vector.Vector
		switch sf.direction {
		case input.Left:
			newPos = vector.Vector{
				X: sf.snake.Pos.X - 32,
				Y: sf.snake.Pos.Y,
			}
		case input.Right:
			newPos = vector.Vector{
				X: sf.snake.Pos.X + 32,
				Y: sf.snake.Pos.Y,
			}
		case input.Up:
			newPos = vector.Vector{
				X: sf.snake.Pos.X,
				Y: sf.snake.Pos.Y - 32,
			}
		case input.Down:
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
				sf.DeathCallback()
				return nil
			}
			s = s.Next
		}

		sf.snake.Move(newPos)
		if sf.snake.Pos == sf.foodPos {
			sf.EatCallback()
			sf.snake.Grow()
			sf.newFood()
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

	var foodDrawOptions ebiten.DrawImageOptions
	foodDrawOptions.GeoM.Translate(sf.foodPos.X, sf.foodPos.Y)
	snakeField.DrawImage(image_manager.Food(), &foodDrawOptions)

	screen.DrawImage(snakeField, opts)
}

func (sf *SnakeField) newFood() {
	newFoodPos := vector.Vector{
		X: float64(rand.Intn(config.FieldWidth/config.TileSize) * config.TileSize),
		Y: float64(rand.Intn(config.FieldHeight/config.TileSize) * config.TileSize),
	}

	s := sf.snake
	for s != nil {
		if s.Pos == newFoodPos {
			fmt.Println("collision")
			newFoodPos = vector.Vector{
				X: float64(rand.Intn(config.FieldWidth/config.TileSize) * config.TileSize),
				Y: float64(rand.Intn(config.FieldHeight/config.TileSize) * config.TileSize),
			}
		}
		s = s.Next
	}
	sf.foodPos = newFoodPos
}

func (sf *SnakeField) SpeedUp() {
	sf.speed++
	sf.mainTicker = time.NewTicker(calcTick(sf.speed))
}

func (sf *SnakeField) GrowSnake() {
	sf.snake.Grow()
}

func calcTick(speed int) time.Duration {
	return (time.Second) / time.Duration(speed)
}
