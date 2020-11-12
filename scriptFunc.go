package main

import (
	"fmt"
	"time"
)

func st(text string, script Script) {
	script.text = substr(text, 4, len(text))
	fmt.Println(script.text)
}

func sc(text []string, pos int, amount int, script Script, player Player) {

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
			reader.player = player
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

func nc(text []string, pos int, amount int, npc NPC, player Player) {

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
			reader.player = player
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
		if player.hp < 0 {
			fmt.Println("You Lost")
			return
		}
		fmt.Println("Attack him")
		fight := readline()
		if fight == "attack" {
			enemyNew := player.fight(enemy)
			enemy.hp = enemyNew.hp
			fmt.Printf("\nEnemy has: %shp\n", enemy.hp)
		}
		fmt.Println("Dodge him")
		ch1 := make(chan string)
		go goOne(ch1)

		select {
		case msg := <-ch1:
			if msg != "dodge" {
				playerNew := enemy.enemyAttack(player)
				player.hp = playerNew.hp
				fmt.Println("Missed hit")
				fmt.Printf("\nYou have: %shp\n", player.hp)
			} else {
				fmt.Println("dodged")
			}
		case <-time.After(time.Second * 6):
			playerNew := enemy.enemyAttack(player)
			player.hp = playerNew.hp
			fmt.Println("Missed hit")
			fmt.Printf("\nYou have: %shp\n", player.hp)
		}

	}
	fmt.Println("You killed the monster")
}

func goOne(ch chan string) {
	time.Sleep(time.Second * 5)
	ch <- readline()
}

/*
func main() {
	fmt.Println("Dodge him")
	ch1 := make(chan string)
	go goOne(ch1)

	select {
	case msg := <-ch1:
		if msg != "dodge" {
			fmt.Println("missed hit")
		}else {
			fmt.Println("dodged")
		}
	case <-time.After(time.Second * 6):
		fmt.Println("missed hit")
	}
}
*/
