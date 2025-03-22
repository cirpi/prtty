// It contains the ASCII escape codes for text decoration
package prtty

import "fmt"

type Style string

var (
	// ASCII escape character
	ESC Style = "\x1B"
	// Text Decorations
	Bold          Style = "\x1B[1m"
	Dim           Style = "\x1B[2m"
	Italic        Style = "\x1B[3m"
	Underline     Style = "\x1B[4m"
	Strikethrough Style = "\x1B[9m"
)

// Commonly used colours (Foreground)
var (
	BlackF   Style = Foreground(0, 0, 0)
	BlueF    Style = Foreground(0, 0, 255)
	RedF     Style = Foreground(255, 0, 0)
	GreenF   Style = Foreground(0, 0, 0)
	WhiteF   Style = Foreground(255, 255, 255)
	MagentaF Style = Foreground(200, 0, 127)
	YellowF  Style = Foreground(255, 255, 102)
	OrangeF  Style = Foreground(255, 153, 51)
)

// Commonly used colours (Background)
var (
	BlackB   Style = Background(0, 0, 0)
	BlueB    Style = Background(0, 0, 255)
	RedB     Style = Background(255, 0, 0)
	GreenB   Style = Background(0, 0, 0)
	WhiteB   Style = Background(255, 255, 255)
	MagentaB Style = Background(200, 0, 127)
	YellowB  Style = Background(255, 255, 102)
	OrangeB  Style = Background(255, 153, 51)
)

// Function to create your custom foreground colour.
func Foreground(r, g, b uint) Style {
	return Style(fmt.Sprintf("\x1b[38;2;%d;%d;%dm", r, g, b))
}

// Function to create you custom background colour.
func Background(r, g, b uint) Style {
	return Style(fmt.Sprintf("\x1b[48;2;%d;%d;%dm", r, g, b))
}
