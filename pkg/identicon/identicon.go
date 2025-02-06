package identicon

import (
	"crypto/md5"
	"errors"
	"image"
	"image/color"
	"image/png"
	"io"
)

var (
	ErrInvalidInput     = errors.New("invalid input: empty string")
	ErrNilWriter        = errors.New("nil writer provided")
	ErrInvalidSize      = errors.New("invalid size: must be positive")
	ErrInvalidGrid      = errors.New("invalid grid size: must be positive odd number")
	ErrInvalidColor     = errors.New("invalid color: nil color provided")
	ErrInvalidGenerator = errors.New("invalid generator: nil generator provided")
	ErrInvalidPadding   = errors.New("invalid padding: must be between 0 and 0.5")
	ErrInvalidMinPoints = errors.New("invalid min points: must be positive")
)

type Identicon struct {
	config *Config
}

func New(opts ...Option) (*Identicon, error) {
	cfg := defaultConfig()

	for _, opt := range opts {
		if err := opt(cfg); err != nil {
			return nil, err
		}
	}

	return &Identicon{
		config: cfg,
	}, nil
}

func (i *Identicon) Generate(input string) (*image.RGBA, error) {
	if input == "" {
		return nil, ErrInvalidInput
	}

	hash := md5.Sum([]byte(input))
	img := i.createImage()

	foreground := i.config.ColorGenerator.Generate(hash)
	i.drawPattern(img, hash, foreground)

	return img, nil
}

func (i *Identicon) SavePNG(img *image.RGBA, w io.Writer) error {
	if w == nil {
		return ErrNilWriter
	}
	return png.Encode(w, img)
}

func (i *Identicon) createImage() *image.RGBA {
	img := image.NewRGBA(image.Rect(0, 0, i.config.Size, i.config.Size))
	for x := 0; x < i.config.Size; x++ {
		for y := 0; y < i.config.Size; y++ {
			img.Set(x, y, i.config.Background)
		}
	}
	return img
}

func (i *Identicon) drawPattern(img *image.RGBA, hash [16]byte, c color.Color) {
	cellSize := i.config.Size / i.config.GridSize
	padding := int(float64(cellSize) * i.config.Padding)
	pointsDrawn := 0

	for x := 0; x < i.config.GridSize/2+1; x++ {
		for y := 0; y < i.config.GridSize; y++ {
			shouldDraw := hash[(x+y*i.config.GridSize)%16]%2 == 0

			if shouldDraw {
				// Draw left side
				i.fillRect(img, x*cellSize+padding, y*cellSize+padding,
					cellSize-padding*2, cellSize-padding*2, c)
				pointsDrawn++

				// Draw mirrored right side (except center column)
				if x != i.config.GridSize/2 {
					i.fillRect(img, (i.config.GridSize-1-x)*cellSize+padding,
						y*cellSize+padding, cellSize-padding*2, cellSize-padding*2, c)
					pointsDrawn++
				}
			}
		}
	}

	// NEW: Second pass - ensure minimum points if needed
	if pointsDrawn < i.config.MinPoints {
		remainingPoints := i.config.MinPoints - pointsDrawn
		centerX := i.config.GridSize / 2

		for y := 0; y < i.config.GridSize && remainingPoints > 0; y++ {
			if hash[y%16]%3 == 0 && !i.hasPoint(img, centerX, y, cellSize, c) {
				i.fillRect(img, centerX*cellSize+padding, y*cellSize+padding,
					cellSize-padding*2, cellSize-padding*2, c)
				remainingPoints--
			}
		}

		if remainingPoints > 0 {
			for y := 0; y < i.config.GridSize && remainingPoints > 0; y++ {
				if !i.hasPoint(img, centerX, y, cellSize, c) {
					i.fillRect(img, centerX*cellSize+padding, y*cellSize+padding,
						cellSize-padding*2, cellSize-padding*2, c)
					remainingPoints--
				}
			}
		}
	}
}

func (i *Identicon) fillRect(img *image.RGBA, x, y, width, height int, c color.Color) {
	for dx := 0; dx < width; dx++ {
		for dy := 0; dy < height; dy++ {
			img.Set(x+dx, y+dy, c)
		}
	}
}

func (i *Identicon) hasPoint(img *image.RGBA, gridX, gridY, cellSize int, c color.Color) bool {
	x := gridX*cellSize + cellSize/2
	y := gridY*cellSize + cellSize/2
	r, g, b, a := img.At(x, y).RGBA()
	pointR, pointG, pointB, pointA := c.RGBA()
	return r == pointR && g == pointG && b == pointB && a == pointA
}
