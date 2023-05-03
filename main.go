package main

import (
	"github.com/df-mc/structure"
	"github.com/eminarican/libabel/library"
	"github.com/eminarican/libabel/library/session"
	"golang.design/x/clipboard"
	"log"
)

func main() {
	if err := clipboard.Init(); err != nil {
		log.Fatalf("couldn't initialize clipboard: %v", err)
	}

	str, err := structure.ReadFile("./assets/library.mcstructure")
	if err != nil {
		log.Fatalf("couldn't read structure file: %v", err)
	}
	gen := library.NewGenerator(str)

	cfg, err := library.DefaultConfig(gen)
	if err != nil {
		log.Fatalf("couldn't create config: %v", err)
	}

	pro, err := session.NewProvider("./players")
	if err != nil {
		log.Fatalf("couldn't create session data provider: %v", err)
	}

	lib := library.New(cfg, pro)
	lib.Start()
}
