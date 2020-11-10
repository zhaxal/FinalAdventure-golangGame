package main

import (
	"bufio"
	"fmt"
	"os"
)

type NamePicker struct {
	playerName string
}

func (n *NamePicker) namePick() {
	fmt.Println("Enter your name")
	n.playerName = readline()
}

func (n *NamePicker) getName() string {
	return n.playerName
}

type Intro struct{}

func (i *Intro) showIntro() {
	fmt.Println("-=Final Adventure=-\nAn adventure game made in Go!")
}

type Greetings struct {
	playerName NamePicker
}

func (g *Greetings) greet() {

}

func readline() string {
	reader := bufio.NewReader(os.Stdin)
	byteName, _, _ := reader.ReadLine()
	return string(byteName)
}

type Singleton interface {
	AddOne() int
}

type singleton struct {
	count int
}

var instance *singleton

func greet() Singleton {
	if instance == nil {
		instance = new(singleton)

		g := NamePicker{}

		g.namePick()

		playerName := g.getName()

		if playerName == "" {
			playerName = "player"
		}

		fmt.Printf("\nAlrighty, %s.\n", playerName)
	}
	return instance
}

func (s *singleton) AddOne() int {
	s.count++
	return s.count
}
