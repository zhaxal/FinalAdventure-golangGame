package main

import (
	"fmt"
	"time"
)

func st(text string, script Script) {
	script.text = substr(text, 4, len(text))
	fmt.Println(script.text)
}

func sc(text []string, pos int, amount int, script Script) {

	script.choice = make([]string, amount)

	fmt.Println("")
	fmt.Println("Your choice: ")

	for i := 0; i < amount; i++ {
		script.choice[i] = text[pos+i+1]
		fmt.Println(script.choice[i])
	}

	fmt.Println("")

	choice := readline()

	fmt.Println("")

	for i := range script.choice {
		if script.choice[i] == choice {
			reader := &Reader{}
			reader.scriptReader(choice + ".txt")
		}
	}
}

func nn(text string, npc NPC) {
	npc.name = substr(text, 4, len(text))
}

func nt(text string, npc NPC) {
	npc.script.text = substr(text, 4, len(text))
}

func nc(text []string, pos int, amount int, npc NPC) {

	npc.script.choice = make([]string, amount)

	fmt.Println("")
	fmt.Println("Your choice: ")

	for i := 0; i < amount; i++ {
		npc.script.choice[i] = text[pos+i+1]
		fmt.Println(npc.script.choice[i])
	}
	fmt.Println("")

	choice := readline()

	fmt.Println("")
	for i := range npc.script.choice {
		if npc.script.choice[i] == choice {
			reader := &Reader{}
			reader.scriptReader(choice + ".txt")
		}
	}
}

func es(text string, enemy Enemy, pos int, player Player) {
	enemy.enemyName = substr(text, 4, len(text))
	enemy.hp = int(text[pos+1])
	enemy.enemyDamage = int(text[pos+2])

	fmt.Println("")
	fmt.Printf("You have monster called: %s\n", enemy.enemyName)

	for 0 < enemy.hp {
		fmt.Println("Attack him")
		fight := readline()
		if fight == "attack" {
			player.fight(enemy)
		}
		fmt.Println("Dodge him")
		dodge := ""

		newtimer := time.NewTimer(5 * time.Second)
		<-newtimer.C

		if dodge != "dodge" {
			enemy.enemyAttack(player)
		}

		dodge = readline()

	}
}

func main() {
	var canceled = make(chan struct{})
	dodge := ""
	select {
	case <-time.After(5 * time.Second):
		// do something for timeout, like change state
		dodge = readline()
	case <-canceled:
		// aborted
		if dodge != "dodge" {
			fmt.Println("udar")
		}
	}

}
