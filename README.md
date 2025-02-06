# Go Identicon Generator

The Go Identicon Generator is a package for creating GitHub-style identicons. It generates unique, deterministic, and visually appealing avatar images based on input strings such as email addresses or usernames.

## Features

This package provides a flexible and efficient solution for generating identicons with several key features:

- Deterministic generation ensures the same input always produces the same image
- Customizable image size and grid dimensions
- Adjustable padding between pattern elements
- Configurable minimum number of pattern points
- Custom color generation support
- Symmetrical pattern generation
- PNG output format support

## Installation

To use this package in your project, install it using Go modules:

```bash
go get github.com/navnitms/go-identicon
```

## Basic Usage

Here's a simple example of generating an identicon:

```go
package main

import (
    "image/color"
    "log"
    "os"

    "github.com/navnitms/go-identicon/pkg/identicon"
)

func main() {
    // Create a new generator with default settings
    generator, err := identicon.New()
    if err != nil {
        log.Fatal(err)
    }

    // Generate an identicon
    img, err := generator.Generate("user@example.com")
    if err != nil {
        log.Fatal(err)
    }

    // Save the generated image
    f, err := os.Create("avatar.png")
    if err != nil {
        log.Fatal(err)
    }
    defer f.Close()

    if err := generator.SavePNG(img, f); err != nil {
        log.Fatal(err)
    }
}
```

## Advanced Configuration

The package supports various customization options through the functional options pattern:

```go
generator, err := identicon.New(
    identicon.WithSize(300),          // Set image size to 300x300 pixels
    identicon.WithGridSize(7),        // Use a 7x7 grid for more complex patterns
    identicon.WithPadding(0.15),      // Add 15% padding between cells
    identicon.WithMinPoints(5),       // Ensure at least 5 points in the pattern
    identicon.WithBackground(color.White), // Set white background
)
```

## Custom Color Generation

You can implement custom color generation by creating a type that satisfies the ColorGenerator interface:

```go
type CustomColorGenerator struct {
    baseColor color.Color
}

func (g *CustomColorGenerator) Generate(hash [16]byte) color.Color {
    r, g, b, _ := g.baseColor.RGBA()
    return color.RGBA{
        R: uint8((r + uint32(hash[0])) / 257),
        G: uint8((g + uint32(hash[1])) / 257),
        B: uint8((b + uint32(hash[2])) / 257),
        A: 255,
    }
}

// Usage
generator, err := identicon.New(
    identicon.WithColorGenerator(&CustomColorGenerator{
        baseColor: color.RGBA{R: 100, G: 150, B: 200, A: 255},
    }),
)
```

## Command Line Usage

The package includes a command-line tool for generating identicons:

```bash
# Build the command-line tool
go build -o identicon ./cmd/identicon

# Generate an identicon
./identicon -input="user@example.com" -output="avatar.png" -size=300 -grid=7 -padding=0.15
```

## Error Handling

The package provides detailed error information for various scenarios:

```go
generator, err := identicon.New(
    identicon.WithSize(-100), // Will return ErrInvalidSize
)
if err != nil {
    switch err {
    case identicon.ErrInvalidSize:
        log.Println("Size must be positive")
    case identicon.ErrInvalidGrid:
        log.Println("Grid size must be positive odd number")
    case identicon.ErrInvalidPadding:
        log.Println("Padding must be between 0 and 0.5")
    default:
        log.Println("Unexpected error:", err)
    }
}
```

## Example Outputs
Here are some sample identicons generated using the package:

![Identicons1](./docs/images/avatar_1.png)
![Identicons2](./docs/images/avatar_2.png)
![Identicons3](./docs/images/avatar_3.png)
![Identicons4](./docs/images/avatar_4.png)

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request. For major changes, please open an issue first to discuss what you would like to change.

## License

This project is licensed under the MIT License - see the LICENSE file for details.

## Acknowledgments

This package was inspired by GitHub's identicon generation system