package main

import (
	"fmt"
)

type Player struct {
	playerName NamePicker
	hp         int
	weapon     Weapon
	inventory  Inventory
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

type Weapon struct {
	weaponName   string
	WeaponDamage int
	isExists     bool
}

func (w *Weapon) Attack(enemy Enemy) Enemy {
	if w.isExists {
		enemy.hp = enemy.hp - w.WeaponDamage
	} else {
		fmt.Println("You don't have a weapon!")
	}
	return enemy
}

type Enemy struct {
	enemyName   string
	enemyDamage int
	hp          int
}

func (e *Enemy) Attack(player Player) Player {
	player.hp = player.hp - e.enemyDamage
	return player
}
