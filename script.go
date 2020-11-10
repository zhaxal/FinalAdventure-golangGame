package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

type Script struct {
	text   string
	choice []string
}

type NPC struct {
	name   string
	script Script
}

type Reader struct {
	filePath string
	script   []string
}

func (r *Reader) readFile() {
	f, err := os.Open(r.filePath)

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {

		if r.script[0] == "" {
			r.script[0] = scanner.Text()
		} else {
			r.script = append(r.script, scanner.Text())
		}

	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func (r *Reader) scriptReader(filepath string) {
	r.filePath = filepath
	r.script = make([]string, 1)
	r.readFile()

	script := Script{}
	npc := NPC{}

	/*
		enemy := Enemy{}

	*/

	for i, s := range r.script {
		switch substr(s, 0, 3) {
		case "-st":
			st(s, script)
		case "-sc":
			amount, err := strconv.Atoi(substr(s, 3, 4))
			sc(r.script, i, amount, script)
			if err != nil {
				fmt.Println(err)
			}
		case "-nn":
			nn(s, npc)
		case "-nt":
			nt(s, npc)
		case "-nc":
			amount, err := strconv.Atoi(substr(s, 3, 4))
			nc(r.script, i, amount, npc)
			if err != nil {
				fmt.Println(err)
			}
		}
	}

}

func substr(input string, start int, length int) string {
	asRunes := []rune(input)

	if start >= len(asRunes) {
		return ""
	}

	if start+length > len(asRunes) {
		length = len(asRunes) - start
	}

	return string(asRunes[start : start+length])
}
