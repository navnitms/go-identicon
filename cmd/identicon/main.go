package main

import (
	"flag"
	"fmt"
	"image/color"
	"log"
	"os"
	"path/filepath"

	"github.com/navnitms/go-identicon/pkg/identicon"
)

func main() {
	var (
		input    string
		output   string
		size     int
		gridSize int
		padding  float64
	)

	flag.StringVar(&input, "input", "", "Input string to generate identicon from")
	flag.StringVar(&output, "output", "identicon.png", "Output file path")
	flag.IntVar(&size, "size", 420, "Size of the identicon in pixels")
	flag.IntVar(&gridSize, "grid", 5, "Grid size (must be odd number)")
	flag.Float64Var(&padding, "padding", 0.1, "Padding between cells (0-0.5)")
	flag.Parse()

	if input == "" {
		fmt.Println("Error: input string is required")
		flag.Usage()
		os.Exit(1)
	}

	// Create generator with provided options
	generator, err := identicon.New(
		identicon.WithSize(size),
		identicon.WithGridSize(gridSize),
		identicon.WithBackground(color.White),
		identicon.WithPadding(padding),
	)
	if err != nil {
		log.Fatal("Failed to create generator:", err)
	}

	// Generate the identicon
	img, err := generator.Generate(input)
	if err != nil {
		log.Fatal("Failed to generate identicon:", err)
	}

	// Ensure output directory exists
	if err := os.MkdirAll(filepath.Dir(output), 0755); err != nil {
		log.Fatal("Failed to create output directory:", err)
	}

	// Create output file
	f, err := os.Create(output)
	if err != nil {
		log.Fatal("Failed to create output file:", err)
	}
	defer f.Close()

	// Save the identicon
	if err := generator.SavePNG(img, f); err != nil {
		log.Fatal("Failed to save identicon:", err)
	}

	fmt.Printf("Successfully generated identicon: %s\n", output)
}
