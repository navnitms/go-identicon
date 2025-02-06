package main

import (
	"log"
	"os"
	"github.com/navnitms/go-identicon/pkg/identicon"
)

func main() {
	generator, err := identicon.New()
	if err != nil {
		log.Fatal(err)
	}
	
	img, err := generator.Generate("example2@email.com")
	if err != nil {
		log.Fatal(err)
	}
	
	f, err := os.Create("avatar.png")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	
	if err := generator.SavePNG(img, f); err != nil {
		log.Fatal(err)
	}
}
