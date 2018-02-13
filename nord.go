package nord

import (
	"fmt"
	"strings"
)

type Color struct {
	r, g, b uint32
}

// RGBA satisfies image.Color interface
func (c Color) RGBA() (r, g, b, a uint32) {
	return c.r, c.g, c.b, 1
}

// Hex outputs as #rrggbb
func (c Color) Hex() string {
	return fmt.Sprintf("#%02x%02x%02x", c.r, c.g, c.b)
}

// String satisfies fmt.Stringer interface
// alias for Hex
func (c Color) String() string {
	return c.Hex()
}

// ParseHex returns a Color from the format #rgb or #rrggbb
func ParseHex(hex string) (Color, error) {
	var r, g, b uint32
	hex = strings.ToLower(hex)
	if len(hex) == 4 {
		_, err := fmt.Sscanf(hex, "#%01x%01x%01x", &r, &g, &b)
		return Color{r * 0x11, g * 0x11, b * 0x11}, err

	} else if len(hex) == 7 {
		_, err := fmt.Sscanf(hex, "#%02x%02x%02x", &r, &g, &b)
		return Color{r, g, b}, err
	}
	return Color{}, fmt.Errorf("ambiguous input color: %v", hex)
}

var (
	Nord0          = Color{0x2e, 0x34, 0x40}
	PolarNightBase = Nord0

	Nord1              = Color{0x3b, 0x42, 0x52}
	PolarNightLighter1 = Nord1

	Nord2              = Color{0x43, 0x4c, 0x5e}
	PolarNightLighter2 = Nord2

	Nord3              = Color{0x4c, 0x56, 0x6a}
	PolarNightLighter3 = Nord3

	Nord4         = Color{0xd8, 0xde, 0xe9}
	SnowStormBase = Nord4

	Nord5             = Color{0xe5, 0xe9, 0xf0}
	SnowStormLighter1 = Nord5

	Nord6             = Color{0xec, 0xef, 0xf4}
	SnowStormLighter2 = Nord6

	Nord7     = Color{0x8f, 0xbc, 0xbb}
	FrostBase = Nord7

	Nord8         = Color{0x88, 0xc0, 0xd0}
	FrostLighter1 = Nord8

	Nord9         = Color{0x81, 0xa1, 0xc1}
	FrostLighter2 = Nord9

	Nord10        = Color{0x5e, 0x81, 0xac}
	FrostLighter3 = Nord10

	Nord11    = Color{0xbf, 0x61, 0x6a}
	AuroraRed = Nord11

	Nord12       = Color{0xd0, 0x87, 0x70}
	AuroraOrange = Nord12

	Nord13       = Color{0xeb, 0xcb, 0x8b}
	AuroraYellow = Nord13

	Nord14      = Color{0xa3, 0xbe, 0x8c}
	AuroraGreen = Nord14

	Nord15       = Color{0xb4, 0x8e, 0xad}
	AuroraPurple = Nord15
)
