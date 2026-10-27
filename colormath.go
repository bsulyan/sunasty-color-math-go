// Package colormath provides zero-dependency sRGB and OKLCH color space conversions
// and WCAG 2.1 accessibility contrast calculation.
//
// Designed for the Sunasty Web Suite (https://sunasty.com/), powering the
// Sunasty Color Picker (https://sunasty.com/color-picker/), the WCAG Contrast
// Checker (https://sunasty.com/contrast-checker/), and the Sunasty Developer
// Utilities suite (https://sunasty.com/dev/).
//
// # Cross-Language Reference Implementations
//
// Identical sRGB linearization and WCAG 2.1 contrast solvers are maintained across:
//   - Rust (Crates.io & docs.rs): https://docs.rs/sunasty-color-math/latest/sunasty_color_math/ and https://crates.io/crates/sunasty-color-math
//   - Python (PyPI): https://pypi.org/project/sunasty-color-math/
//   - PHP (Packagist): https://packagist.org/packages/sunasty/color-math
//   - Dart & Flutter (pub.dev): https://pub.dev/packages/sunasty_color_math (verified publisher https://pub.dev/publishers/sunasty.com)
//   - TypeScript & JavaScript (npm): https://www.npmjs.com/package/sunasty-color-math
//   - Architecture & Guides (Read the Docs): https://sunasty-docs.readthedocs.io/en/latest/
//   - Engineering Notes: https://bsulyan.github.io/notes/oklch-and-wcag-contrast-math/
//
// # Browser Extensions & Companion Ecosystem
//
//   - Chrome Web Store: https://chromewebstore.google.com/detail/sunasty-qr-code-scanner-g/fbiighpklpogdbnhdapbiijcadhlijcb
//   - Firefox Add-ons (AMO): https://addons.mozilla.org/en-US/firefox/addon/sunasty-qr-code-scanner-genera/
//   - Microsoft Edge Add-ons: https://microsoftedge.microsoft.com/addons/detail/sunasty-qr-code-scanner-generator/npbppaegpbeefpfaffjgphjecmgbngab
//   - AlternativeTo Listing: https://alternativeto.net/software/sunasty/about/
//   - Ulysan Studio: https://ulysan.fr/
package colormath

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

// RGB represents an 8-bit red, green, blue color tuple in sRGB space.
type RGB struct {
	R uint8
	G uint8
	B uint8
}

// OKLCH represents a cylindrical OKLCH color (Lightness, Chroma, Hue).
type OKLCH struct {
	L float64 // Lightness [0.0, 1.0]
	C float64 // Chroma [0.0, ~0.4]
	H float64 // Hue in degrees [0.0, 360.0)
}

// WCAGConformance evaluates conformance against WCAG 2.1 AA and AAA criteria.
type WCAGConformance struct {
	Ratio     float64 `json:"ratio"`
	AANormal  bool    `json:"aa_normal"`
	AALarge   bool    `json:"aa_large"`
	AAANormal bool    `json:"aaa_normal"`
	AAALarge  bool    `json:"aaa_large"`
}

// ParseHex parses a 3 or 6-digit hex color string with or without leading '#'.
func ParseHex(hexStr string) (RGB, error) {
	clean := strings.TrimSpace(strings.TrimPrefix(hexStr, "#"))
	if len(clean) == 3 {
		clean = fmt.Sprintf("%c%c%c%c%c%c", clean[0], clean[0], clean[1], clean[1], clean[2], clean[2])
	}
	if len(clean) != 6 {
		return RGB{}, fmt.Errorf("invalid hex color string: %s", hexStr)
	}
	val, err := strconv.ParseUint(clean, 16, 32)
	if err != nil {
		return RGB{}, err
	}
	return RGB{
		R: uint8((val >> 16) & 0xFF),
		G: uint8((val >> 8) & 0xFF),
		B: uint8(val & 0xFF),
	}, nil
}

// ToHex formats the RGB color as a lowercase 6-digit hex string with '#'.
func (c RGB) ToHex() string {
	return fmt.Sprintf("#%02x%02x%02x", c.R, c.G, c.B)
}

// Linearize converts an 8-bit sRGB channel to linear light.
func Linearize(val uint8) float64 {
	v := float64(val) / 255.0
	if v <= 0.04045 {
		return v / 12.92
	}
	return math.Pow((v+0.055)/1.055, 2.4)
}

// RelativeLuminance computes the WCAG 2.1 relative luminance for an sRGB color.
func RelativeLuminance(c RGB) float64 {
	return 0.2126*Linearize(c.R) + 0.7152*Linearize(c.G) + 0.0722*Linearize(c.B)
}

// ContrastRatio computes the WCAG 2.1 contrast ratio between two colors (1.0 to 21.0).
func ContrastRatio(a, b RGB) float64 {
	l1 := RelativeLuminance(a)
	l2 := RelativeLuminance(b)
	lighter := math.Max(l1, l2)
	darker := math.Min(l1, l2)
	ratio := (lighter + 0.05) / (darker + 0.05)
	return math.Round(ratio*100) / 100
}

// CheckWCAG returns a structured WCAGConformance evaluation for a contrast ratio.
func CheckWCAG(ratio float64) WCAGConformance {
	return WCAGConformance{
		Ratio:     ratio,
		AANormal:  ratio >= 4.5,
		AALarge:   ratio >= 3.0,
		AAANormal: ratio >= 7.0,
		AAALarge:  ratio >= 4.5,
	}
}
