package nord

import (
	"fmt"
	"strings"
)

// Color holds the internal representation of a color
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
	// Base component color of "Polar Night".
	//
	// Used for texts, backgrounds, carets and structuring characters like curly- and square brackets.
	Nord0          = Color{0x2e, 0x34, 0x40}
	PolarNightBase = Nord0

	// Lighter shade color of the base component color.
	//
	// Used as a lighter background color for UI elements like status bars.
	Nord1              = Color{0x3b, 0x42, 0x52}
	PolarNightLighter1 = Nord1

	// Lighter shade color of the base component color.
	//
	// Used as line highlighting in the editor.
	// In the UI scope it may be used as selection- and hightlight color.
	Nord2              = Color{0x43, 0x4c, 0x5e}
	PolarNightLighter2 = Nord2

	// Lighter shade color of the base component color.
	//
	// Used for comments, invisibles, indent- and wrap guide marker.
	// In the UI scope used as pseudoclass color for disabled elements.
	Nord3              = Color{0x4c, 0x56, 0x6a}
	PolarNightLighter3 = Nord3

	// Base component color of "Snow Storm".
	//
	// Main color for text, variables, constants and attributes.
	// In the UI scope used as semi-light background depending on the theme shading design.
	Nord4         = Color{0xd8, 0xde, 0xe9}
	SnowStormBase = Nord4

	// Lighter shade color of the base component color.
	//
	// Used as a lighter background color for UI elements like status bars.
	// Used as semi-light background depending on the theme shading design.
	Nord5             = Color{0xe5, 0xe9, 0xf0}
	SnowStormLighter1 = Nord5

	// Lighter shade color of the base component color.
	//
	// Used for punctuations, carets and structuring characters like curly- and square brackets.
	// In the UI scope used as background, selection- and hightlight color depending on the theme shading design.
	Nord6             = Color{0xec, 0xef, 0xf4}
	SnowStormLighter2 = Nord6

	// Bluish core color.
	//
	// Used for classes, types and documentation tags.
	Nord7     = Color{0x8f, 0xbc, 0xbb}
	FrostBase = Nord7

	// Bluish core color.
	//
	// Represents the accent color of the color palette.
	// Main color for primary UI elements and methods/functions.
	//
	// Can be used for
	//   - Markup quotes
	//   - Markup link URLs
	Nord8         = Color{0x88, 0xc0, 0xd0}
	FrostLighter1 = Nord8

	// Bluish core color.
	//
	// Used for language-specific syntactic/reserved support characters and keywords, operators, tags, units and
	// punctuations like (semi)colons,commas and braces.
	Nord9         = Color{0x81, 0xa1, 0xc1}
	FrostLighter2 = Nord9

	// Bluish core color.
	//
	// Used for markup doctypes, import/include/require statements, pre-processor statements and at-rules (`@`).
	Nord10        = Color{0x5e, 0x81, 0xac}
	FrostLighter3 = Nord10

	// Colorful component color.
	//
	// Used for errors, git/diff deletion and linter marker.
	Nord11    = Color{0xbf, 0x61, 0x6a}
	AuroraRed = Nord11

	// Colorful component color.
	//
	// Used for annotations.
	Nord12       = Color{0xd0, 0x87, 0x70}
	AuroraOrange = Nord12

	// Colorful component color.
	//
	// Used for escape characters, regular expressions and markup entities.
	// In the UI scope used for warnings and git/diff renamings.
	Nord13       = Color{0xeb, 0xcb, 0x8b}
	AuroraYellow = Nord13

	// Colorful component color.
	//
	// Main color for strings and attribute values.
	// In the UI scope used for git/diff additions and success visualizations.
	Nord14      = Color{0xa3, 0xbe, 0x8c}
	AuroraGreen = Nord14

	// Colorful component color.
	//
	// Used for numbers.
	Nord15       = Color{0xb4, 0x8e, 0xad}
	AuroraPurple = Nord15
)
