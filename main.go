package main

import (
	"github.com/df-mc/structure"
	"github.com/eminarican/libabel/library"
	"log"
)

func main() {
	str, err := structure.ReadFile("./assets/library.mcstructure")
	if err != nil {
		log.Fatalf("couldn't read structure file: %v", err)
	}
	gen := library.NewGenerator(str)

	cfg, err := library.DefaultConfig(gen)
	if err != nil {
		log.Fatalf("couldn't create config: %v", err)
	}

	lib := library.New(cfg)
	lib.Start()
}
