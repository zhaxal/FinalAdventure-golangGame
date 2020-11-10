package main

import "fmt"

type Singleton interface {
	AddOne() int
}

type singleton struct {
	count int
}

var instance *singleton
var raceList []string

func pickRace(playerName string) (Singleton, Player) {
	player := Player{}
	if instance == nil {
		raceList = []string{"elf", "orc", "human", "druid"}
		fmt.Println("Choose your race: ")
		fmt.Println("")
		for i := range raceList {
			fmt.Println(raceList[i])
		}
		fmt.Println("")
		race := readline()

		player = getRace(race, playerName)

	}

	return instance, player
}

func (s *singleton) AddOne() int {
	s.count++
	return s.count
}
