package main

type elf struct {
	Player
}

func newElf(name string) *Player {
	return &Player{
		playerName:   NamePicker{name},
		hp:           50,
		attackDamage: 20,
		weapon:       &BowBehavior{},
		inventory: Inventory{
			slots: 6,
		},
	}
}

type orc struct {
	Player
}

func newOrc(name string) *Player {
	return &Player{
		playerName:   NamePicker{name},
		hp:           100,
		attackDamage: 40,
		weapon:       &AxeBehavior{},
		inventory: Inventory{
			slots: 2,
		},
	}
}

type human struct {
	Player
}

func newHuman(name string) *Player {
	return &Player{
		playerName:   NamePicker{name},
		hp:           70,
		attackDamage: 30,
		weapon:       &SwordBehavior{},
		inventory: Inventory{
			slots: 4,
		},
	}
}

type druid struct {
	Player
}

func newDruid(name string) *Player {
	return &Player{
		playerName:   NamePicker{name},
		hp:           80,
		attackDamage: 25,
		weapon:       &WandBehavior{},
		inventory: Inventory{
			slots: 3,
		},
	}
}

func getRace(playerRace string, name string) Player {
	switch playerRace {
	case "elf":
		return *newElf(name)
	case "orc":
		return *newOrc(name)
	case "human":
		return *newHuman(name)
	case "druid":
		return *newDruid(name)
	}
	return Player{}
}
