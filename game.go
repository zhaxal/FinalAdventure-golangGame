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

type Intro struct{}

func (i *Intro) showIntro() {
	fmt.Println("-=Final Adventure=-\\nAn adventure game made in Go!")
}

type Greetings struct {
	playerName NamePicker
}

func (g *Greetings) greet() {

	g.playerName.namePick()

	if g.playerName.playerName == "" {
		g.playerName.playerName = "player"
	}

	fmt.Printf("\nAlrighty, %s.\n", g.playerName.playerName)
}

type GameFacade struct {
	intro     *Intro
	greetings *Greetings
}

func NewGameFacade() *GameFacade {
	return &GameFacade{new(Intro), new(Greetings)}
}

func (g *GameFacade) start() {
	g.intro.showIntro()
	g.greetings.greet()
}

func main() {
	game := NewGameFacade()
	game.start()
}
