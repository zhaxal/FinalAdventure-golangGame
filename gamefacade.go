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
	g.intro.showIntro()
	g.greetings.greet()
	g.reader.scriptReader()
}

/*
func main() {
	game := NewGameFacade()
	game.start()
}
*/
