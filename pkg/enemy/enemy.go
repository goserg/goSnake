package enemy

import (
	"fmt"
	"time"
)

type Enemy struct {
	name  string
	maxHP int
	HP    int

	cooldown   time.Duration
	nextAttack time.Time

	attackCallback func()
	deathCallback  func()
}

func New(attackCallback func(),
	deathCallback func()) *Enemy {
	var enemy Enemy
	enemy.name = "rat"
	enemy.maxHP = 100
	enemy.HP = 100
	enemy.cooldown = time.Second * 2
	enemy.nextAttack = time.Now().Add(enemy.cooldown)
	enemy.attackCallback = attackCallback
	enemy.deathCallback = deathCallback

	return &enemy
}

func (e *Enemy) Update() {
	if time.Now().After(e.nextAttack) {
		e.nextAttack = time.Now().Add(e.cooldown)
		fmt.Printf("%s attack\n", e.name)
		e.attackCallback()
	}
}

func (e *Enemy) Damage(dmg int) {
	e.HP -= dmg
	if e.HP <= 0 {
		e.deathCallback()
	}
}
