package colormath

import (
	"math"
	"testing"
)

func TestParseHex(t *testing.T) {
	tests := []struct {
		hex      string
		expected RGB
	}{
		{"#ffffff", RGB{255, 255, 255}},
		{"000000", RGB{0, 0, 0}},
		{"#0f766e", RGB{15, 118, 110}},
		{"#fff", RGB{255, 255, 255}},
	}

	for _, tt := range tests {
		got, err := ParseHex(tt.hex)
		if err != nil {
			t.Fatalf("ParseHex(%q) unexpected error: %v", tt.hex, err)
		}
		if got != tt.expected {
			t.Errorf("ParseHex(%q) = %v; want %v", tt.hex, got, tt.expected)
		}
	}
}

func TestContrastRatio(t *testing.T) {
	black := RGB{0, 0, 0}
	white := RGB{255, 255, 255}

	ratio := ContrastRatio(white, black)
	if math.Abs(ratio-21.0) > 0.1 {
		t.Errorf("ContrastRatio(white, black) = %v; want ~21.0", ratio)
	}

	wcag := CheckWCAG(ratio)
	if !wcag.AANormal || !wcag.AAANormal {
		t.Errorf("CheckWCAG(21.0) failed: %+v", wcag)
	}
}
