package game

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"goSnake/pkg/config"
	"goSnake/pkg/enemy"
	"goSnake/pkg/input"
	"goSnake/pkg/item"
	snakeField "goSnake/pkg/snake_field"
	"goSnake/pkg/text"
	"goSnake/pkg/ui"
	"golang.org/x/image/colornames"
	"time"
)

type Game struct {
	snakeField *snakeField.SnakeField
	enemy      *enemy.Enemy

	isDebugVisible bool

	input *input.Handler

	ui *ui.UI
}

func (g *Game) IsDisposed() bool {
	return false
}

func New() *Game {
	var g Game

	g.ui = ui.New()
	g.ui.EventStartPressed.Connect(&g, g.OnStartButtonPressed)

	inputHandler := input.NewHandler()
	g.input = inputHandler

	g.isDebugVisible = true

	g.snakeField = snakeField.New(g.input)
	g.snakeField.EventDeath.Connect(&g, g.onDeath)
	g.snakeField.EventEat.Connect(&g, g.OnSnakeEatEvent)

	g.enemy = enemy.New()
	g.enemy.EventAttack.Connect(&g, g.onEnemyAttack)
	g.enemy.EventDeath.Connect(&g, g.onEnemyDeath)
	g.enemy.EventTakeDamage.Connect(&g, g.onEnemyTakeDamage)
	return &g
}

func (g *Game) Update() error {
	g.ui.Update()

	g.input.Update()
	text.Update()
	if g.input.IsActionJustPressed(input.Debug) {
		g.isDebugVisible = !g.isDebugVisible
	}
	if g.input.IsActionJustPressed(input.Pause) {
		g.ui.ToggleMenu()
		g.snakeField.Toggle()
		g.enemy.Toggle()
	}

	if err := g.snakeField.Update(); err != nil {
		return err
	}

	g.enemy.Update()
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
	enemyDrawingOptions.GeoM.Translate(700, 100)
	g.enemy.Draw(screen, &enemyDrawingOptions)

	text.Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideHeight
}

func (g *Game) onDeath(data snakeField.EventSnakeDeathData) {
	g.snakeField = snakeField.New(g.input)
	g.snakeField.EventDeath.Connect(g, g.onDeath)
	g.snakeField.EventEat.Connect(g, g.OnSnakeEatEvent)
	g.enemy = enemy.New()
	g.enemy.EventAttack.Connect(g, g.onEnemyAttack)
	g.enemy.EventDeath.Connect(g, g.onEnemyDeath)
	g.enemy.EventTakeDamage.Connect(g, g.onEnemyTakeDamage)

	g.ui.ShowMenu()

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
		g.snakeField.SpawnRock()
	case enemy.AttackTypeGrow:
		g.snakeField.GrowSnake()
	}
}

func (g *Game) onEnemyDeath(data enemy.EventDeathData) {
	g.snakeField = snakeField.New(g.input)
	g.snakeField.EventDeath.Connect(g, g.onDeath)
	g.snakeField.EventEat.Connect(g, g.OnSnakeEatEvent)
	g.enemy = enemy.New()
	g.enemy.EventAttack.Connect(g, g.onEnemyAttack)
	g.enemy.EventDeath.Connect(g, g.onEnemyDeath)
	g.enemy.EventTakeDamage.Connect(g, g.onEnemyTakeDamage)

	g.ui.ShowMenu()

	text.New("YOU WIN", 200, 200,
		text.WithColor(colornames.Greenyellow),
		text.WithSize(100),
		text.WithFadeout(),
		text.WithMove(0, -0.5),
		text.WithLifespan(time.Second),
	)
}

func (g *Game) OnSnakeEatEvent(arg snakeField.EventEatData) {
	switch arg.Type {
	case item.TypeSword:
		g.enemy.Damage(10)
	case item.TypeRock:
		g.onDeath(snakeField.EventSnakeDeathData{})
	}
}

func (g *Game) OnStartButtonPressed(data struct{}) {
	fmt.Println("start")
	g.snakeField.Start()
	g.enemy.Start()
	g.ui.HideMenu()
}

func (g *Game) onEnemyTakeDamage(data enemy.EventTakeDamageData) {
	text.New(fmt.Sprintf("-%d", data.Value), 700, 100,
		text.WithColor(colornames.Red),
		text.WithSize(14),
		text.WithMove(0, -0.5),
		text.WithLifespan(time.Second),
	)
}
