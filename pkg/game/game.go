package game

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"goSnake/pkg/config"
	"goSnake/pkg/enemy"
	"goSnake/pkg/input"
	snakeField "goSnake/pkg/snake_field"
	"goSnake/pkg/text"
	"golang.org/x/image/colornames"
	"time"
)

type Game struct {
	snakeField     *snakeField.SnakeField
	isDebugVisible bool

	enemy *enemy.Enemy
}

func (g *Game) IsDisposed() bool {
	return false
}

func New() *Game {
	var g Game

	g.isDebugVisible = true

	g.snakeField = snakeField.New()
	g.snakeField.EventDeath.Connect(&g, g.onDeath)
	g.snakeField.EventEat.Connect(&g, g.OnSnakeEatEvent)

	g.enemy = enemy.New()
	g.enemy.EventAttack.Connect(&g, g.onEnemyAttack)
	g.enemy.EventDeath.Connect(&g, g.onEnemyDeath)
	return &g
}

func (g *Game) Update() error {
	input.Update()
	text.Update()
	if input.IsF1Pressed() {
		g.isDebugVisible = !g.isDebugVisible
	}

	if err := g.snakeField.Update(); err != nil {
		return err
	}

	g.enemy.Update()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	var snakeFieldDrawingOptions ebiten.DrawImageOptions
	snakeFieldDrawingOptions.GeoM.Translate(config.FieldLeft, config.FieldTop)
	g.snakeField.Draw(screen, &snakeFieldDrawingOptions)

	if g.isDebugVisible {
		ebitenutil.DebugPrint(screen, fmt.Sprintf("FPS: %d", int(ebiten.ActualFPS())))
		ebitenutil.DebugPrintAt(screen, fmt.Sprintf("TPS: %d", int(ebiten.ActualTPS())), 0, 10)
	}

	text.Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideHeight
}

func (g *Game) onDeath(data snakeField.EventSnakeDeathData) {
	g.snakeField = snakeField.New()
	g.snakeField.EventDeath.Connect(g, g.onDeath)
	g.snakeField.EventEat.Connect(g, g.OnSnakeEatEvent)
	g.enemy = enemy.New()
	g.enemy.EventAttack.Connect(g, g.onEnemyAttack)
	g.enemy.EventDeath.Connect(g, g.onEnemyDeath)

	text.New("YOU DIED", 200, 200,
		text.WithColor(colornames.Red),
		text.WithSize(100),
		text.WithFadeout(),
		text.WithMove(0, -0.5),
		text.WithLifespan(time.Second),
	)
}

func (g *Game) onEnemyAttack(data enemy.EventAttackData) {
	g.snakeField.GrowSnake()
}

func (g *Game) onEnemyDeath(data enemy.EventDeathData) {
	g.snakeField = snakeField.New()
	g.snakeField.EventDeath.Connect(g, g.onDeath)
	g.snakeField.EventEat.Connect(g, g.OnSnakeEatEvent)
	g.enemy = enemy.New()
	g.enemy.EventAttack.Connect(g, g.onEnemyAttack)
	g.enemy.EventDeath.Connect(g, g.onEnemyDeath)

	text.New("YOU WIN", 200, 200,
		text.WithColor(colornames.Greenyellow),
		text.WithSize(100),
		text.WithFadeout(),
		text.WithMove(0, -0.5),
		text.WithLifespan(time.Second),
	)
}

func (g *Game) OnSnakeEatEvent(arg snakeField.EventEatData) {
	fmt.Println("event eat", arg.Name)
	g.enemy.Damage(10)
}
