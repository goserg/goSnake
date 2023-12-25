package enemy

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"goSnake/pkg/ui"
	"goSnake/resource"

	//"github.com/hajimehoshi/ebiten/v2"
	"goSnake/pkg/engine/signal"
	"time"
)

type Enemy struct {
	isRunning bool

	name  string
	maxHP int
	HP    int

	img       *ebiten.Image
	healthBar *ui.HealthBar

	cooldown   time.Duration
	nextAttack time.Time

	EventDeath  signal.Event[EventDeathData]
	EventAttack signal.Event[EventAttackData]
}

type EventDeathData struct {
}
type EventAttackData struct {
}

func New() *Enemy {
	var enemy Enemy
	enemy.name = "rat"
	enemy.maxHP = 100
	enemy.HP = 100
	enemy.img = resource.Image(resource.ImageSkeleton)
	enemy.healthBar = ui.NewHealthBar()
	enemy.cooldown = time.Second * 2
	enemy.nextAttack = time.Now().Add(enemy.cooldown)

	return &enemy
}

func (e *Enemy) Update() {
	if !e.isRunning {
		return
	}
	if time.Now().After(e.nextAttack) {
		e.nextAttack = time.Now().Add(e.cooldown)
		fmt.Printf("%s attack\n", e.name)
		e.EventAttack.Emit(EventAttackData{})
	}
}

func (e *Enemy) Damage(dmg int) {
	e.HP -= dmg
	e.healthBar.Set(e.HP * 100 / e.maxHP)
	if e.HP <= 0 {
		e.EventDeath.Emit(EventDeathData{})
	}
}

func (e *Enemy) Start() {
	e.isRunning = true
}

func (e *Enemy) Toggle() {
	e.isRunning = !e.isRunning
}

func (e *Enemy) Draw(screen *ebiten.Image, drawingOptions *ebiten.DrawImageOptions) {
	screen.DrawImage(e.img, drawingOptions)

	barImg := ebiten.NewImage(64, 3)
	e.healthBar.Draw(barImg)
	drawingOptions.GeoM.Translate(-16, -5)

	screen.DrawImage(barImg, drawingOptions)
}
