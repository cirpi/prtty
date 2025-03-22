// To create a new style
package prtty

import (
	"fmt"
	"strings"
)

// struct to hold all the decorations
type decorator struct {
	styles []Style
}

// constructor to create a new decorator
func NewDecorator(styles ...Style) *decorator {
	d := &decorator{}
	for _, style := range styles {
		d.styles = append(d.styles, style)
	}
	return d
}

// apply the styles to the given input string.
func (d *decorator) String(str string) string {
	var styleString strings.Builder
	for _, style := range d.styles {
		styleString.WriteString(string(style))
	}
	return fmt.Sprintf("%s%s%s", styleString.String(), str, Reset)
}
