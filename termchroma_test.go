package termchroma

import (
	"strings"
	"testing"
)

// TestHexToRGB checks parsing of valid and invalid hex codes.
func TestHexToRGB(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		wantR   int
		wantG   int
		wantB   int
		wantErr bool
	}{
		{"Full #000000", "#000000", 0, 0, 0, false},
		{"Full #ffffff", "#ffffff", 255, 255, 255, false},
		{"Short #abc => #aabbcc", "#abc", 170, 187, 204, false},
		{"Short #123 => #112233", "#123", 17, 34, 51, false},
		{"No # sign", "ff0000", 255, 0, 0, false},
		{"Invalid length", "#1234", 0, 0, 0, true},
		{"Invalid chars", "#gggggg", 0, 0, 0, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r, g, b, err := HexToRGB(tt.input)
			if (err != nil) != tt.wantErr {
				t.Fatalf("HexToRGB(%q) error = %v, wantErr = %v", tt.input, err, tt.wantErr)
			}
			if !tt.wantErr {
				if r != tt.wantR || g != tt.wantG || b != tt.wantB {
					t.Fatalf("HexToRGB(%q) => (%d,%d,%d), want (%d,%d,%d)",
						tt.input, r, g, b, tt.wantR, tt.wantG, tt.wantB)
				}
			}
		})
	}
}

// TestANSIForeground ensures we return the correct escape code for foreground colors.
func TestANSIForeground(t *testing.T) {
	code, err := ANSIForeground("#ff0000")
	if err != nil {
		t.Fatalf("ANSIForeground error: %v", err)
	}
	if !strings.HasPrefix(code, "\033[38;2;255;0;0m") {
		t.Errorf("ANSIForeground(#ff0000) => %q, want prefix %q", code, "\033[38;2;255;0;0m")
	}
}

// TestANSIBackground ensures we return the correct escape code for background colors.
func TestANSIBackground(t *testing.T) {
	code, err := ANSIBackground("#00ff00")
	if err != nil {
		t.Fatalf("ANSIBackground error: %v", err)
	}
	if !strings.HasPrefix(code, "\033[48;2;0;255;0m") {
		t.Errorf("ANSIBackground(#00ff00) => %q, want prefix %q", code, "\033[48;2;0;255;0m")
	}
}
