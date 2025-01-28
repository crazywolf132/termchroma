package termchroma

import (
	"errors"
	"strconv"
	"strings"
)

// Some helpful default ansi codes
const (
	Bold      = "\033[1m"
	Dim       = "\033[2m"
	Underline = "\033[4m"
	Blink     = "\033[5m"
	Reverse   = "\033[7m"
	Hidden    = "\033[8m"
	Reset     = "\033[0m"
)

// HexToRGB converts a 3- or 6-digit hex color (#RRGGBB or #RGB) into integer RGB components (0-255).
// It ignores a leading "#" if present. Returns an error if the string is invalid.
func HexToRGB(hex string) (int, int, int, error) {
	hex = strings.TrimPrefix(hex, "#")
	hex = strings.ToLower(hex)

	// Validate length first
	if len(hex) != 3 && len(hex) != 6 {
		return 0, 0, 0, errors.New("invalid hex color: must be 3 or 6 hex digits")
	}

	// Validate that all characters are valid hex digits
	for _, c := range hex {
		if (c < '0' || c > '9') && (c < 'a' || c > 'f') {
			return 0, 0, 0, errors.New("invalid hex color: contains non-hex characters")
		}
	}

	// Expand 3-digit #RGB to #RRGGBB
	if len(hex) == 3 {
		hex = string([]byte{
			hex[0], hex[0],
			hex[1], hex[1],
			hex[2], hex[2],
		})
	}

	// Parse each two-digit component
	rVal, err := strconv.ParseInt(hex[0:2], 16, 32)
	if err != nil {
		return 0, 0, 0, errors.New("invalid red component")
	}
	gVal, err := strconv.ParseInt(hex[2:4], 16, 32)
	if err != nil {
		return 0, 0, 0, errors.New("invalid green component")
	}
	bVal, err := strconv.ParseInt(hex[4:6], 16, 32)
	if err != nil {
		return 0, 0, 0, errors.New("invalid blue component")
	}

	return int(rVal), int(gVal), int(bVal), nil
}

// ANSIForeground returns the ANSI escape code to set the terminal's foreground
// (text) color to the provided hex color code (#RRGGBB or #RGB).
func ANSIForeground(hex string) (string, error) {
	r, g, b, err := HexToRGB(hex)
	if err != nil {
		return "", err
	}
	return "\033[38;2;" +
		strconv.Itoa(r) + ";" +
		strconv.Itoa(g) + ";" +
		strconv.Itoa(b) + "m", nil
}

// ANSIBackground returns the ANSI escape code to set the terminal's background
// color to the provided hex color code (#RRGGBB or #RGB).
func ANSIBackground(hex string) (string, error) {
	r, g, b, err := HexToRGB(hex)
	if err != nil {
		return "", err
	}
	return "\033[48;2;" +
		strconv.Itoa(r) + ";" +
		strconv.Itoa(g) + ";" +
		strconv.Itoa(b) + "m", nil
}
