# termchroma

A lightweight Go library for converting hex color codes to ANSI escape sequences for terminal output.

## Installation

```bash
go get github.com/crazywolf132/termchroma
```

## Features

- Converts hex color codes to RGB values
- Supports both 6-digit (#RRGGBB) and 3-digit (#RGB) hex codes
- Generates ANSI escape sequences for both foreground and background colors
- No external dependencies
- Thoroughly tested

## Usage

```go
package main

import (
    "fmt"
    "github.com/crazywolf132/termchroma"
)

func main() {
    // Set foreground (text) color using hex code
    fgCode, _ := termchroma.ANSIForeground("#ff0000")
    fmt.Printf("%sRed Text%s\n", fgCode, termchroma.Reset)

    // Set background color using hex code
    bgCode, _ := termchroma.ANSIBackground("#00ff00")
    fmt.Printf("%sGreen Background%s\n", bgCode, termchroma.Reset)

    // Convert hex color to RGB values
    r, g, b, _ := termchroma.HexToRGB("#0000ff")
    fmt.Printf("RGB values: (%d, %d, %d)\n", r, g, b)
}
```

## API

### `HexToRGB(hex string) (r, g, b int, err error)`

Converts a hex color code to RGB values. Supports both 6-digit (#RRGGBB) and 3-digit (#RGB) formats.
The function accepts hex codes with or without the leading '#' character.

Returns an error if:
- The input length is not 3 or 6 characters (excluding '#')
- The input contains non-hex characters

### `ANSIForeground(hex string) (string, error)`

Converts a hex color code to an ANSI escape sequence for setting the terminal's foreground (text) color.
Returns the escape sequence as a string, which can be printed to the terminal.

### `ANSIBackground(hex string) (string, error)`

Converts a hex color code to an ANSI escape sequence for setting the terminal's background color.
Returns the escape sequence as a string, which can be printed to the terminal.

### Constants

- `Reset`: ANSI escape code for resetting colors to default (`"\033[0m"`)

## License

MIT