package main

import (
	"bufio"
	"log"
	"os"
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
