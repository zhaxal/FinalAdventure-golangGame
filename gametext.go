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

type Procedure interface {
	Request(flag bool)
}

type Intro struct {
	next Procedure
}

func (i *Intro) Request(flag bool) {
	fmt.Println("-=Final Adventure=-\nAn adventure game made in Go!")
	if flag {
		i.next.Request(flag)
	}
}

type Greetings struct {
	next   Procedure
	picker NamePicker
}

func (g *Greetings) Request(flag bool) {
	g.picker.namePick()

	playerName := g.picker.getName()

	if playerName == "" {
		playerName = "player"
	}

	fmt.Printf("\nAlrighty, %s.\n", playerName)
}

func readline() string {
	reader := bufio.NewReader(os.Stdin)
	byteName, _, _ := reader.ReadLine()
	return string(byteName)
}
