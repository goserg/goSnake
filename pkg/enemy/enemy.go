package enemy

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"goSnake/pkg/engine/signal"
	"goSnake/pkg/ui"
	"goSnake/resource"
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
	nextAttack time.Duration

	EventDeath      signal.Event[EventDeathData]
	EventTakeDamage signal.Event[EventTakeDamageData]
	EventAttack     signal.Event[EventAttackData]
}

type EventDeathData struct {
}
type EventTakeDamageData struct {
	Value int
}

type AttackType int

const (
	AttackTypeRock AttackType = iota + 1
	AttackTypeGrow
	AttackTypeSpeedUp
	AttackTypeSlowDown
)

type EventAttackData struct {
	AttackType AttackType
}

func New() *Enemy {
	var enemy Enemy
	enemy.name = "rat"
	enemy.maxHP = 100
	enemy.HP = 100
	enemy.img = resource.Image(resource.ImageSkeleton)
	enemy.healthBar = ui.NewHealthBar()
	enemy.cooldown = time.Second * 2
	enemy.nextAttack = enemy.cooldown

	return &enemy
}

func (e *Enemy) Update(delta time.Duration) {
	if !e.isRunning {
		return
	}
	e.nextAttack -= delta
	if e.nextAttack < 0 {
		e.nextAttack = e.cooldown
		fmt.Printf("%s attack\n", e.name)
		e.EventAttack.Emit(EventAttackData{
			AttackType: AttackTypeRock,
		})
	}
}

func (e *Enemy) Damage(dmg int) {
	e.HP -= dmg
	e.healthBar.Set(e.HP * 100 / e.maxHP)
	e.EventTakeDamage.Emit(EventTakeDamageData{Value: dmg})
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
