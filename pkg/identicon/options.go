package identicon

import (
	"image/color"

	icolor "github.com/navnitms/go-identicon/pkg/identicon/color"
)

// Config holds all the configuration options for the identicon generator
type Config struct {
	// Size represents the width and height of the generated image in pixels
	Size int

	// GridSize determines the number of cells in the pattern grid
	// Must be an odd number to maintain symmetry
	GridSize int

	// Background specifies the background color of the generated image
	Background color.Color

	// ColorGenerator defines how colors are generated for the pattern
	ColorGenerator icolor.ColorGenerator

	// Padding specifies the space between cells as a ratio of cell size
	// Must be between 0 and 0.5
	Padding float64

	// MinPoints specifies the minimum number of points in the polygon
	// Must be greater than 0
	MinPoints int
}

type Option func(*Config) error

func defaultConfig() *Config {
	return &Config{
		Size:           420,
		GridSize:       5,
		Background:     color.White,
		ColorGenerator: &icolor.DefaultGenerator{},
		Padding:        0.1,
		MinPoints:      4,
	}
}

// WithSize sets the size of the generated identicon
// The size must be a positive number
func WithSize(size int) Option {
	return func(c *Config) error {
		if size <= 0 {
			return ErrInvalidSize
		}
		c.Size = size
		return nil
	}
}

// WithGridSize sets the grid size for the pattern
// The grid size must be a positive odd number
func WithGridSize(size int) Option {
	return func(c *Config) error {
		if size <= 0 || size%2 == 0 {
			return ErrInvalidGrid
		}
		c.GridSize = size
		return nil
	}
}

// WithBackground sets the background color
// The color must not be nil
func WithBackground(c color.Color) Option {
	return func(cfg *Config) error {
		if c == nil {
			return ErrInvalidColor
		}
		cfg.Background = c
		return nil
	}
}

// WithColorGenerator sets a custom color generator
// The generator must not be nil
func WithColorGenerator(g icolor.ColorGenerator) Option {
	return func(c *Config) error {
		if g == nil {
			return ErrInvalidGenerator
		}
		c.ColorGenerator = g
		return nil
	}
}

// WithPadding sets the padding between cells
// The padding must be between 0 and 0.5
func WithPadding(padding float64) Option {
	return func(c *Config) error {
		if padding < 0 || padding >= 0.5 {
			return ErrInvalidPadding
		}
		c.Padding = padding
		return nil
	}
}

// WithMinPoints sets the minimum number of points in the polygon
// The minimum points must be greater than 0
func WithMinPoints(points int) Option {
    return func(c *Config) error {
        if points < 0 {
            return ErrInvalidMinPoints
        }
        c.MinPoints = points
        return nil
    }
}
