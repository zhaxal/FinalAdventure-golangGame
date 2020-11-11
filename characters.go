package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Player struct {
	playerName   NamePicker
	hp           int
	attackDamage int
	weapon       WeaponBehavior
	inventory    Inventory
}

type Inventory struct {
	items []string
	slots int
}

func (i *Inventory) addItem(item string) {
	isAdded := false
	for r := range i.items {
		if i.items[r] == "" {
			i.items[r] = item
			isAdded = true
			return
		}
	}
	if isAdded {
		fmt.Println("Item added successfully")
	} else {
		fmt.Println("Not enough space")
	}

}

func (i *Inventory) removeItem(item string) {
	isRemoved := false
	for r := range i.items {
		if i.items[r] == item {
			i.items[r] = ""
			isRemoved = true
			return
		}
	}
	if isRemoved {
		fmt.Println("Item removed successfully")
	} else {
		fmt.Println("You dont have such item")
	}
}

func (player *Player) fight(enemy Enemy) {
	player.weapon.attack(enemy, player.attackDamage)
}

type WeaponBehavior interface {
	attack(enemy Enemy, attackDamage int) *Enemy
}

func (player *Player) setWeaponBehavior(weaponToSet WeaponBehavior) {
	player.weapon = weaponToSet
}

type SwordBehavior struct {
}

func (sb SwordBehavior) attack(enemy Enemy, attackDamage int) *Enemy {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	chance := r1.Intn(4)
	if chance == 1 {
		enemy.hp = enemy.hp - int(float64(attackDamage)*1.5)
	} else {
		enemy.hp = enemy.hp - attackDamage
	}
	return &enemy
}

type WandBehavior struct {
}

func (wb WandBehavior) attack(enemy Enemy, attackDamage int) *Enemy {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	chance := r1.Intn(10)
	if chance == 1 {
		enemy.hp = enemy.hp - 2*attackDamage
	} else {
		enemy.hp = enemy.hp - attackDamage
	}
	return &enemy
}

type BowBehavior struct {
}

func (bb BowBehavior) attack(enemy Enemy, attackDamage int) *Enemy {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	chance := r1.Intn(5)
	if chance == 1 {
		enemy.hp = enemy.hp - 3*attackDamage
	} else if chance == 2 {

	} else {
		enemy.hp = enemy.hp - attackDamage
	}
	return &enemy
}

type AxeBehavior struct {
}

func (ab AxeBehavior) attack(enemy Enemy, attackDamage int) *Enemy {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	chance := r1.Intn(8)
	if chance == 1 {
		enemy.hp = enemy.hp - int(float64(attackDamage)*2.5)
	} else {
		enemy.hp = enemy.hp - attackDamage
	}
	return &enemy
}

type Enemy struct {
	enemyName   string
	enemyDamage int
	hp          int
}

func (e *Enemy) enemyAttack(player Player) *Player {
	player.hp = player.hp - e.enemyDamage
	return &player
}
