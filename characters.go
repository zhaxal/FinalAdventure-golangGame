package main

type Player struct {
	playerName   NamePicker
	hp           int
	attackDamage int
	weapon       WeaponBehavior
}

func (player *Player) fight(enemy Enemy) *Enemy {
	return player.weapon.attack(enemy, player.attackDamage)
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
