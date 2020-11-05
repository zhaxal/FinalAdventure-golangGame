package main

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
