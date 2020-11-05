package main

type Player struct {
	playerName NamePicker
	hp         int
}

type Weapon struct {
	weaponName   string
	WeaponDamage int
	isCapable    bool
}

func (w *Weapon) Attack(enemy Enemy) Enemy {
	if w.isCapable {
		enemy.hp = enemy.hp - w.WeaponDamage
	}
	return enemy
}

type Enemy struct {
	enemyName   string
	enemyDamage int
	hp          int
}
