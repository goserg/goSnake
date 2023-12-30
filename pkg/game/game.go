package game

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"goSnake/pkg/config"
	"goSnake/pkg/enemy"
	"goSnake/pkg/input"
	"goSnake/pkg/inventory"
	"goSnake/pkg/item"
	snakeField "goSnake/pkg/snake_field"
	"goSnake/pkg/text"
	"goSnake/pkg/ui"
	"goSnake/pkg/utils/vector"
	"goSnake/resource"
	"golang.org/x/image/colornames"
	"strconv"
	"time"
)

type Game struct {
	inventory  *inventory.Inventory
	snakeField *snakeField.SnakeField
	enemy      *enemy.Enemy

	isDebugVisible bool

	input *input.Handler

	ui *ui.UI

	lastFrameTime time.Time
}

func (g *Game) IsDisposed() bool {
	return false
}

func New() *Game {
	var g Game

	initGame(&g)
	return &g
}

func initGame(g *Game) {
	g.inventory = &inventory.Inventory{Items: []item.Type{item.TypeFood, item.TypeFood, item.TypePotionSpeedUp, item.TypePotionSpeedUp, item.TypePotionSpeedUp, item.TypePotionSpeedUp}}

	g.ui = ui.New()
	g.ui.EventStartPressed.Connect(g, g.OnStartButtonPressed)

	inputHandler := input.NewHandler()
	g.input = inputHandler

	g.isDebugVisible = true

	g.snakeField = snakeField.New(g.input)
	g.snakeField.EventDeath.Connect(g, g.onDeath)
	g.snakeField.EventEat.Connect(g, g.onSnakeEatEvent)
	g.snakeField.EventItemSpawned.Connect(g, g.onItemSpawned)
	for _, itemType := range g.inventory.Items {
		g.snakeField.SpawnItem(itemType)
	}

	g.enemy = enemy.New()
	g.enemy.EventAttack.Connect(g, g.onEnemyAttack)
	g.enemy.EventDeath.Connect(g, g.onEnemyDeath)
	g.enemy.EventTakeDamage.Connect(g, g.onEnemyTakeDamage)

	g.lastFrameTime = time.Now()
}

func (g *Game) Update() error {
	delta := time.Since(g.lastFrameTime)
	g.lastFrameTime = time.Now()

	g.ui.Update(delta)

	g.input.Update(delta)
	text.Update(delta)
	if g.input.IsActionJustPressed(input.Debug) {
		g.isDebugVisible = !g.isDebugVisible
	}
	if g.input.IsActionJustPressed(input.Pause) {
		g.ui.ToggleMenu()
		g.snakeField.Toggle()
		g.enemy.Toggle()
	}

	if err := g.snakeField.Update(delta); err != nil {
		return err
	}

	g.enemy.Update(delta)
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.ui.Draw(screen)

	var snakeFieldDrawingOptions ebiten.DrawImageOptions
	snakeFieldDrawingOptions.GeoM.Translate(config.FieldLeft, config.FieldTop)
	g.snakeField.Draw(screen, &snakeFieldDrawingOptions)

	if g.isDebugVisible {
		ebitenutil.DebugPrint(screen, fmt.Sprintf("FPS: %d", int(ebiten.ActualFPS())))
		ebitenutil.DebugPrintAt(screen, fmt.Sprintf("TPS: %d", int(ebiten.ActualTPS())), 0, 10)
	}

	var enemyDrawingOptions ebiten.DrawImageOptions
	enemyDrawingOptions.GeoM.Translate(900, 100)
	g.enemy.Draw(screen, &enemyDrawingOptions)

	text.Draw(screen)

	if m != nil {
		m.Draw(screen)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideHeight
}

func (g *Game) onDeath(data snakeField.EventSnakeDeathData) {
	initGame(g)

	text.New("YOU DIED", 200, 200,
		text.WithColor(colornames.Red),
		text.WithSize(100),
		text.WithFadeout(),
		text.WithMove(0, -0.5),
		text.WithLifespan(time.Second),
	)
}

func (g *Game) onEnemyAttack(data enemy.EventAttackData) {
	switch data.AttackType {
	case enemy.AttackTypeRock:
		g.snakeField.SpawnItem(item.TypeRock)
	case enemy.AttackTypeGrow:
		g.snakeField.GrowSnake()
	}
}

func (g *Game) onEnemyDeath(data enemy.EventDeathData) {
	initGame(g)

	text.New("YOU WIN", 200, 200,
		text.WithColor(colornames.Greenyellow),
		text.WithSize(100),
		text.WithFadeout(),
		text.WithMove(0, -0.5),
		text.WithLifespan(time.Second),
	)
}

func (g *Game) onSnakeEatEvent(arg snakeField.EventEatData) {
	switch arg.Type {
	case item.TypeSword:
		g.enemy.Damage(5)

		m = &Missile{
			img: resource.Image(resource.ImageSword),
			startPos: vector.Vector{
				X: arg.Pos.X + config.FieldLeft,
				Y: arg.Pos.Y + config.FieldTop,
			},
			startTime: time.Now(),
			targetPos: vector.Vector{
				X: 900 + config.TileSize,
				Y: 100 + config.TileSize,
			},
			speed: time.Second / 5,
		}
	case item.TypeRock:
		g.onDeath(snakeField.EventSnakeDeathData{})
	case item.TypeFood:
		g.snakeField.GrowSnake()
	case item.TypePotionSpeedUp:
		g.snakeField.SpeedUp()
		text.New("speed up", arg.Pos.X, arg.Pos.Y,
			text.WithColor(colornames.Blue),
			text.WithSize(16),
			text.WithMove(0, -0.5),
			text.WithLifespan(time.Second),
		)
	}
}

type Missile struct {
	img       *ebiten.Image
	startPos  vector.Vector
	startTime time.Time
	targetPos vector.Vector
	speed     time.Duration
}

func (m2 *Missile) Draw(screen *ebiten.Image) {
	nowTime := time.Now().Sub(m2.startTime)
	movedPart := float64(nowTime) / float64(m2.speed)
	prevPos := m2.startPos
	var visualPos vector.Vector
	visualPos.X = prevPos.X + (m2.targetPos.X-prevPos.X)*movedPart
	visualPos.Y = prevPos.Y + (m2.targetPos.Y-prevPos.Y)*movedPart

	var snakeDrawOptions ebiten.DrawImageOptions
	snakeDrawOptions.GeoM.Rotate(movedPart * 4)
	snakeDrawOptions.GeoM.Translate(visualPos.X, visualPos.Y)
	screen.DrawImage(m2.img, &snakeDrawOptions)
}

var m *Missile

func (g *Game) OnStartButtonPressed(data struct{}) {
	fmt.Println("start")
	g.snakeField.Start()
	g.enemy.Start()
	g.ui.HideMenu()
}

func (g *Game) onEnemyTakeDamage(data enemy.EventTakeDamageData) {
	text.New(strconv.Itoa(data.Value), 912, 75,
		text.WithColor(colornames.Yellow),
		text.WithSize(16),
		text.WithMove(0, -0.5),
		text.WithLifespan(time.Second),
	)
}

func (g *Game) onItemSpawned(data snakeField.EventItemSpawnedData) {
	switch data.ItemType {
	case item.TypeRock:
		text.New("rock!", data.Pos.X+config.FieldLeft, data.Pos.Y+config.FieldTop,
			text.WithLifespan(time.Second),
			text.WithSize(16),
			text.WithColor(colornames.Red),
			text.WithMove(0, -0.5),
		)
	}
}
