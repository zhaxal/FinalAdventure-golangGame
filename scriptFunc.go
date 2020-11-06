package main

func st(text string, script Script) {
	script.text = substr(text, 4, len(text))
}

func sc(text []string, pos int, amount int, script Script) {
	for i := 0; i < amount; i++ {
		script.choice[i] = text[pos+i+1]
	}
}

func nn(text string, npc NPC) {
	npc.name = substr(text, 4, len(text))
}

func nt(text string, npc NPC) {
	npc.script.text = substr(text, 4, len(text))
}

func nc(text []string, pos int, amount int, npc NPC) {
	for i := 0; i < amount; i++ {
		npc.script.choice[i] = text[pos+i+1]
	}
}
