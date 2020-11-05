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
	reader := bufio.NewReader(os.Stdin)
	byteName, _, _ := reader.ReadLine()
	n.playerName = string(byteName)
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

	g.playerName.namePick()

	playerName := g.playerName.getName()

	if playerName == "" {
		playerName = "player"
	}

	fmt.Printf("\nAlrighty, %s.\n", playerName)
}
