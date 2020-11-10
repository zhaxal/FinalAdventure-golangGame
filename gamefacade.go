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
	greet()
	g.reader.scriptReader("script.txt")
}

func main() {
	game := NewGameFacade()
	game.start()
}
