package main

type Script struct {
	text   string
	choice []string
}

type NPC struct {
	name   string
	script Script
}

type Reader struct{}

func (r *Reader) Read() {

}
