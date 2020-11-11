package main

type GameFacade struct {
	intro     *Intro
	greetings *Greetings
	reader    *Reader
}

func NewGameFacade() *GameFacade {
	return &GameFacade{new(Intro), new(Greetings), new(Reader)}
}

func (g *GameFacade) start() {
	greet := &Intro{new(Greetings)}
	greet.Request(true)

	_, player := pickRace(g.greetings.picker.playerName)

	g.reader.player = player
	g.reader.scriptReader("script.txt")
}

/*
func main() {
	game := NewGameFacade()
	game.start()
}
*/
