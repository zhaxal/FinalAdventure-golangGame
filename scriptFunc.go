package main

import "fmt"

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

	choice := readline()

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

	choice := readline()

	for i := range npc.script.choice {
		if npc.script.choice[i] == choice {
			reader := &Reader{}
			reader.scriptReader(choice + ".txt")
		}
	}
}
