package main

type elf struct {
	Player
}

func newElf(name string) *Player {
	return &Player{
		playerName: NamePicker{name},
		hp:         50,
		weapon:     Weapon{"bow", 20, true},
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
		playerName: NamePicker{name},
		hp:         100,
		weapon:     Weapon{"axe", 40, true},
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
		playerName: NamePicker{name},
		hp:         70,
		weapon:     Weapon{"sword", 30, true},
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
		playerName: NamePicker{name},
		hp:         80,
		weapon:     Weapon{"wand", 25, true},
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
