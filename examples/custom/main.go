package main

import (
	"fmt"
	"image/color"
	"log"
	"os"

	"github.com/navnitms/go-identicon/pkg/identicon"
)

type CustomColorGenerator struct {
	baseColor color.Color
}

func (gen *CustomColorGenerator) Generate(hash [16]byte) color.Color {
	r, g, b, _ := gen.baseColor.RGBA()

	return color.RGBA{
		R: uint8((r + uint32(hash[0])) / 257),
		G: uint8((g + uint32(hash[1])) / 257),
		B: uint8((b + uint32(hash[2])) / 257),
		A: 255,
	}
}

func main() {
	// Create a custom color generator with a base color
	customGenerator := &CustomColorGenerator{
		baseColor: color.RGBA{R: 100, G: 150, B: 200, A: 255},
	}

	// Initialize generator with custom options
	generator, err := identicon.New(
		identicon.WithSize(300),
		identicon.WithGridSize(5),
		identicon.WithBackground(color.White),
		identicon.WithColorGenerator(customGenerator),
		identicon.WithPadding(0.15),
		identicon.WithMinPoints(4),
	)
	if err != nil {
		log.Fatal("Failed to create generator:", err)
	}

	// Generate multiple identicons with different inputs
	inputs := []string{
		"user1@example.com",
		"user2@example.com",
		"user3@example.com",
	}

	for i, input := range inputs {
		img, err := generator.Generate(input)
		if err != nil {
			log.Printf("Failed to generate identicon for %s: %v", input, err)
			continue
		}

		filename := fmt.Sprintf("custom_avatar_%d.png", i+1)
		f, err := os.Create(filename)
		if err != nil {
			log.Printf("Failed to create file %s: %v", filename, err)
			continue
		}

		if err := generator.SavePNG(img, f); err != nil {
			log.Printf("Failed to save identicon to %s: %v", filename, err)
		}

		f.Close()
		log.Printf("Generated identicon: %s", filename)
	}
}
