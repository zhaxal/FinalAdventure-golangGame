package main

import (
	"fmt"
	"math/rand"
	"time"
)

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
		fmt.Println("Critical damage!")
	} else {
		enemy.hp = enemy.hp - attackDamage
		fmt.Println("Successful attack")
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
		fmt.Println("Critical damage!")
	} else {
		enemy.hp = enemy.hp - attackDamage
		fmt.Println("Successful attack")
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
		fmt.Println("Critical damage!")
	} else if chance == 2 {
		fmt.Println("You missed!")
	} else {
		enemy.hp = enemy.hp - attackDamage
		fmt.Println("Successful attack")
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
		fmt.Println("Critical damage!")
	} else {
		enemy.hp = enemy.hp - attackDamage
		fmt.Println("Successful attack")
	}
	return &enemy
}
