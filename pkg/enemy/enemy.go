package enemy

import (
	"fmt"
	"goSnake/pkg/engine/signal"
	"time"
)

type Enemy struct {
	name  string
	maxHP int
	HP    int

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
	enemy.cooldown = time.Second * 2
	enemy.nextAttack = time.Now().Add(enemy.cooldown)

	return &enemy
}

func (e *Enemy) Update() {
	if time.Now().After(e.nextAttack) {
		e.nextAttack = time.Now().Add(e.cooldown)
		fmt.Printf("%s attack\n", e.name)
		e.EventAttack.Emit(EventAttackData{})
	}
}

func (e *Enemy) Damage(dmg int) {
	e.HP -= dmg
	if e.HP <= 0 {
		e.EventDeath.Emit(EventDeathData{})
	}
}
